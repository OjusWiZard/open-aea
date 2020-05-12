# -*- coding: utf-8 -*-
# ------------------------------------------------------------------------------
#
#   Copyright 2018-2019 Fetch.AI Limited
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.
#
# ------------------------------------------------------------------------------

"""
This module contains the classes required for FIPA dialogue management.

- DialogueLabel: The dialogue label class acts as an identifier for dialogues.
- Dialogue: The dialogue class maintains state of a dialogue and manages it.
- Dialogues: The dialogues class keeps track of all dialogues.
"""

from enum import Enum
from typing import Dict, List, Optional, Tuple, Union, cast

from aea.helpers.dialogue.base import Dialogue, DialogueLabel, Dialogues
from aea.mail.base import Address
from aea.protocols.base import Message

from packages.fetchai.protocols.fipa.message import FipaMessage

VALID_PREVIOUS_PERFORMATIVES = {
    FipaMessage.Performative.CFP: [None],
    FipaMessage.Performative.PROPOSE: [FipaMessage.Performative.CFP],
    FipaMessage.Performative.ACCEPT: [FipaMessage.Performative.PROPOSE],
    FipaMessage.Performative.ACCEPT_W_INFORM: [FipaMessage.Performative.PROPOSE],
    FipaMessage.Performative.MATCH_ACCEPT: [
        FipaMessage.Performative.ACCEPT,
        FipaMessage.Performative.ACCEPT_W_INFORM,
    ],
    FipaMessage.Performative.MATCH_ACCEPT_W_INFORM: [
        FipaMessage.Performative.ACCEPT,
        FipaMessage.Performative.ACCEPT_W_INFORM,
    ],
    FipaMessage.Performative.INFORM: [
        FipaMessage.Performative.MATCH_ACCEPT,
        FipaMessage.Performative.MATCH_ACCEPT_W_INFORM,
        FipaMessage.Performative.INFORM,
    ],
    FipaMessage.Performative.DECLINE: [
        FipaMessage.Performative.CFP,
        FipaMessage.Performative.PROPOSE,
        FipaMessage.Performative.ACCEPT,
        FipaMessage.Performative.ACCEPT_W_INFORM,
    ],
}  # type: Dict[FipaMessage.Performative, List[Union[None, FipaMessage.Performative]]]


class FipaDialogue(Dialogue):
    """The FIPA dialogue class maintains state of a dialogue and manages it."""

    STARTING_MESSAGE_ID = 1
    STARTING_TARGET = 0

    class EndState(Enum):
        """This class defines the end states of a dialogue."""

        SUCCESSFUL = 0
        DECLINED_CFP = 1
        DECLINED_PROPOSE = 2
        DECLINED_ACCEPT = 3

    class AgentRole(Enum):
        """This class defines the agent's role in the dialogue."""

        SELLER = "seller"
        BUYER = "buyer"

    def __init__(
        self, dialogue_label: DialogueLabel, is_seller: bool, **kwargs
    ) -> None:
        """
        Initialize a dialogue label.

        :param dialogue_label: the identifier of the dialogue
        :param is_seller: indicates whether the agent associated with the dialogue is a seller or buyer

        :return: None
        """
        Dialogue.__init__(self, dialogue_label=dialogue_label)
        self._is_seller = is_seller
        self._role = (
            FipaDialogue.AgentRole.SELLER if is_seller else FipaDialogue.AgentRole.BUYER
        )

    @property
    def is_seller(self) -> bool:
        """Check whether the agent acts as the seller in this dialogue."""
        return self._is_seller

    @property
    def role(self) -> "FipaDialogue.AgentRole":
        """Get role of agent in dialogue."""
        return self._role

    def is_valid_next_message(self, fipa_msg: Message) -> bool:
        """
        Check whether this is a valid next message in the dialogue.

        :return: True if yes, False otherwise.
        """
        fipa_msg = cast(FipaMessage, fipa_msg)
        this_message_id = fipa_msg.message_id
        this_target = fipa_msg.target
        this_performative = fipa_msg.performative
        last_outgoing_message = cast(FipaMessage, self.last_outgoing_message)
        if last_outgoing_message is None:
            result = (
                this_message_id == FipaDialogue.STARTING_MESSAGE_ID
                and this_target == FipaDialogue.STARTING_TARGET
                and this_performative == FipaMessage.Performative.CFP
            )
        else:
            last_message_id = last_outgoing_message.message_id
            last_target = last_outgoing_message.target
            last_performative = last_outgoing_message.performative
            result = (
                this_message_id == last_message_id + 1
                and this_target == last_target + 1
                and last_performative in VALID_PREVIOUS_PERFORMATIVES[this_performative]
            )
        return result

    def assign_final_dialogue_label(self, final_dialogue_label: DialogueLabel) -> None:
        """
        Assign the final dialogue label.

        :param final_dialogue_label: the final dialogue label
        :return: None
        """
        assert (
            self.dialogue_label.dialogue_starter_reference
            == final_dialogue_label.dialogue_starter_reference
        )
        assert self.dialogue_label.dialogue_responder_reference == ""
        assert final_dialogue_label.dialogue_responder_reference != ""
        assert (
            self.dialogue_label.dialogue_opponent_addr
            == final_dialogue_label.dialogue_opponent_addr
        )
        assert (
            self.dialogue_label.dialogue_starter_addr
            == final_dialogue_label.dialogue_starter_addr
        )
        self._dialogue_label = final_dialogue_label


class FipaDialogueStats(object):
    """Class to handle statistics on the negotiation."""

    def __init__(self) -> None:
        """Initialize a StatsManager."""
        self._self_initiated = {
            FipaDialogue.EndState.SUCCESSFUL: 0,
            FipaDialogue.EndState.DECLINED_CFP: 0,
            FipaDialogue.EndState.DECLINED_PROPOSE: 0,
            FipaDialogue.EndState.DECLINED_ACCEPT: 0,
        }  # type: Dict[FipaDialogue.EndState, int]
        self._other_initiated = {
            FipaDialogue.EndState.SUCCESSFUL: 0,
            FipaDialogue.EndState.DECLINED_CFP: 0,
            FipaDialogue.EndState.DECLINED_PROPOSE: 0,
            FipaDialogue.EndState.DECLINED_ACCEPT: 0,
        }  # type: Dict[FipaDialogue.EndState, int]

    @property
    def self_initiated(self) -> Dict[FipaDialogue.EndState, int]:
        """Get the stats dictionary on self initiated dialogues."""
        return self._self_initiated

    @property
    def other_initiated(self) -> Dict[FipaDialogue.EndState, int]:
        """Get the stats dictionary on other initiated dialogues."""
        return self._other_initiated

    def add_dialogue_endstate(
        self, end_state: FipaDialogue.EndState, is_self_initiated: bool
    ) -> None:
        """
        Add dialogue endstate stats.

        :param end_state: the end state of the dialogue
        :param is_self_initiated: whether the dialogue is initiated by the agent or the opponent

        :return: None
        """
        if is_self_initiated:
            self._self_initiated[end_state] += 1
        else:
            self._other_initiated[end_state] += 1


class FipaDialogues(Dialogues):
    """The FIPA dialogues class keeps track of all dialogues."""

    def __init__(self) -> None:
        """
        Initialize dialogues.

        :return: None
        """
        Dialogues.__init__(self)
        self._initiated_dialogues = {}  # type: Dict[DialogueLabel, FipaDialogue]
        self._dialogues_as_seller = {}  # type: Dict[DialogueLabel, FipaDialogue]
        self._dialogues_as_buyer = {}  # type: Dict[DialogueLabel, FipaDialogue]
        self._dialogue_stats = FipaDialogueStats()

    @property
    def dialogues_as_seller(self) -> Dict[DialogueLabel, FipaDialogue]:
        """Get dictionary of dialogues in which the agent acts as a seller."""
        return self._dialogues_as_seller

    @property
    def dialogues_as_buyer(self) -> Dict[DialogueLabel, FipaDialogue]:
        """Get dictionary of dialogues in which the agent acts as a buyer."""
        return self._dialogues_as_buyer

    @property
    def dialogue_stats(self) -> FipaDialogueStats:
        """Get the dialogue statistics."""
        return self._dialogue_stats

    def is_permitted_for_new_dialogue(self, fipa_msg: Message) -> bool:
        """
        Check whether a fipa message is permitted for a new dialogue.

        That is, the message has to
        - be a CFP, and
        - have the correct msg id and message target
        - have msg counterparty set.

        :param fipa_msg: the fipa message

        :return: a boolean indicating whether the message is permitted for a new dialogue
        """
        empty_dialogue = self.empty_dialogue()
        result = empty_dialogue.is_valid_next_message(fipa_msg)
        return result

    def is_belonging_to_registered_dialogue(
        self, fipa_msg: Message, agent_addr: Address
    ) -> bool:
        """
        Check whether an agent message is part of a registered dialogue.

        :param fipa_msg: the fipa message
        :param agent_addr: the address of the agent

        :return: boolean indicating whether the message belongs to a registered dialogue
        """
        fipa_msg = cast(FipaMessage, fipa_msg)
        dialogue_reference = fipa_msg.dialogue_reference
        alt_dialogue_reference = (dialogue_reference[0], "")
        self_initiated_dialogue_label = DialogueLabel(
            dialogue_reference, fipa_msg.counterparty, agent_addr
        )
        alt_self_initiated_dialogue_label = DialogueLabel(
            alt_dialogue_reference, fipa_msg.counterparty, agent_addr
        )
        other_initiated_dialogue_label = DialogueLabel(
            dialogue_reference, fipa_msg.counterparty, fipa_msg.counterparty
        )
        result = False
        if other_initiated_dialogue_label in self.dialogues:
            other_initiated_dialogue = cast(
                FipaDialogue, self.dialogues[other_initiated_dialogue_label]
            )
            result = other_initiated_dialogue.is_valid_next_message(fipa_msg)
        if self_initiated_dialogue_label in self.dialogues:
            self_initiated_dialogue = cast(
                FipaDialogue, self.dialogues[self_initiated_dialogue_label]
            )
            result = self_initiated_dialogue.is_valid_next_message(fipa_msg)
        if alt_self_initiated_dialogue_label in self._initiated_dialogues:
            self_initiated_dialogue = cast(
                FipaDialogue,
                self._initiated_dialogues[alt_self_initiated_dialogue_label],
            )
            result = self_initiated_dialogue.is_valid_next_message(fipa_msg)
            if result:
                self._initiated_dialogues.pop(alt_self_initiated_dialogue_label)
                final_dialogue_label = DialogueLabel(
                    dialogue_reference,
                    alt_self_initiated_dialogue_label.dialogue_opponent_addr,
                    alt_self_initiated_dialogue_label.dialogue_starter_addr,
                )
                self_initiated_dialogue.assign_final_dialogue_label(
                    final_dialogue_label
                )
                self._add(self_initiated_dialogue)
        return result

    def get_dialogue(self, fipa_msg: Message, agent_addr: Address) -> Dialogue:
        """
        Retrieve the dialogue the fipa_message belongs to.

        :param fipa_msg: the fipa message
        :param agent_addr: the address of the agent

        :return: the dialogue
        """
        result = None
        fipa_msg = cast(FipaMessage, fipa_msg)
        dialogue_reference = fipa_msg.dialogue_reference
        self_initiated_dialogue_label = DialogueLabel(
            dialogue_reference, fipa_msg.counterparty, agent_addr
        )
        other_initiated_dialogue_label = DialogueLabel(
            dialogue_reference, fipa_msg.counterparty, fipa_msg.counterparty
        )
        if other_initiated_dialogue_label in self.dialogues:
            other_initiated_dialogue = cast(
                FipaDialogue, self.dialogues[other_initiated_dialogue_label]
            )
            result = other_initiated_dialogue
        if self_initiated_dialogue_label in self.dialogues:
            self_initiated_dialogue = cast(
                FipaDialogue, self.dialogues[self_initiated_dialogue_label]
            )
            result = self_initiated_dialogue
        if result is None:
            raise ValueError("Should have found dialogue.")
        return result

    def update(
        self,
        message: Message,
        is_self_message: bool,
        agent_addr: Address,
        is_seller: bool,
    ) -> Optional[Dialogue]:
        """
        Update the state of dialogues with a new message.

        If the message is for a new dialogue, a new dialogue is created with 'message' as its first message and returned.
        If the message is addressed to an existing dialogue, the dialogue is retrieved, extended with this message and returned.
        If there are any errors, e.g. the message dialogue reference does not exists, the message is invalid w.r.t. the dialogue, return None.

        :param message: a new message
        :param agent_addr: the agent address
        :return: the new or existing dialogue the message is intended for, or None in case of any errors.
        """
        fipa_msg = cast(FipaMessage, message)
        dialogue_reference = fipa_msg.dialogue_reference

        # new dialogue by other
        if (
            dialogue_reference[0] != ""
            and dialogue_reference[1] == ""
            and not is_self_message
        ):
            dialogue = self.create_opponent_initiated(
                message.counterparty, fipa_msg.dialogue_reference, is_seller
            )
            dialogue.incoming_extend(message)
        # new dialogue by self
        elif (
            dialogue_reference[0] != ""
            and dialogue_reference[1] == ""
            and is_self_message
        ):
            dialogue = self.create_self_initiated(
                message.counterparty, agent_addr, is_seller
            )
            dialogue.outgoing_extend(message)
        # existing dialogue
        else:
            dialogue = self.get_dialogue(message, agent_addr)
            if is_self_message:
                dialogue.outgoing_extend(message)
            else:
                dialogue.incoming_extend(message)
        return dialogue

    def create_self_initiated(
        self,
        dialogue_opponent_addr: Address,
        dialogue_starter_addr: Address,
        is_seller: bool,
    ) -> Dialogue:
        """
        Create a self initiated dialogue.

        :param dialogue_opponent_addr: the pbk of the agent with which the dialogue is kept.
        :param dialogue_starter_addr: the pbk of the agent which started the dialogue
        :param is_seller: boolean indicating the agent role

        :return: the created dialogue.
        """
        dialogue_reference = (str(self._next_dialogue_nonce()), "")
        dialogue_label = DialogueLabel(
            dialogue_reference, dialogue_opponent_addr, dialogue_starter_addr
        )
        dialogue = FipaDialogue(dialogue_label, is_seller)
        self._initiated_dialogues.update({dialogue_label: dialogue})
        return dialogue

    def create_opponent_initiated(
        self,
        dialogue_opponent_addr: Address,
        dialogue_reference: Tuple[str, str],
        is_seller: bool,
    ) -> Dialogue:
        """
        Save an opponent initiated dialogue.

        :param dialogue_opponent_addr: the address of the agent with which the dialogue is kept.
        :param dialogue_reference: the reference of the dialogue.
        :param is_seller: keeps track if the counterparty is a seller.
        :return: the created dialogue
        """
        assert (
            dialogue_reference[0] != "" and dialogue_reference[1] == ""
        ), "Cannot initiate dialogue with preassigned dialogue_responder_reference!"
        new_dialogue_reference = (
            dialogue_reference[0],
            str(self._next_dialogue_nonce()),
        )
        dialogue_label = DialogueLabel(
            new_dialogue_reference, dialogue_opponent_addr, dialogue_opponent_addr
        )
        result = self._create(dialogue_label, is_seller)
        return result

    def _create(self, dialogue_label: DialogueLabel, is_seller: bool) -> FipaDialogue:
        """
        Create a dialogue.

        :param dialogue_label: the dialogue label
        :param is_seller: boolean indicating the agent role

        :return: the created dialogue
        """
        assert dialogue_label not in self.dialogues
        dialogue = FipaDialogue(dialogue_label, is_seller)
        if is_seller:
            assert dialogue_label not in self.dialogues_as_seller
            self._dialogues_as_seller.update({dialogue_label: dialogue})
        else:
            assert dialogue_label not in self.dialogues_as_buyer
            self._dialogues_as_buyer.update({dialogue_label: dialogue})
        self.dialogues.update({dialogue_label: dialogue})
        return dialogue

    def _add(self, dialogue: FipaDialogue) -> None:
        """
        Add a dialogue.

        :param dialogue: the dialogue

        :return: None
        """
        assert dialogue.dialogue_label not in self.dialogues
        if dialogue.is_seller:
            assert dialogue.dialogue_label not in self.dialogues_as_seller
            self._dialogues_as_seller.update({dialogue.dialogue_label: dialogue})
        else:
            assert dialogue.dialogue_label not in self.dialogues_as_buyer
            self._dialogues_as_buyer.update({dialogue.dialogue_label: dialogue})
        self.dialogues.update({dialogue.dialogue_label: dialogue})
