#!/bin/bash -e

minutes_to_start=$1
participants=$2
mode=$3
registry=$4
acn_node=$5
identifier=$(python -c "import uuid; print(uuid.uuid4().hex)")
tac_name=tac_$identifier
PEER=""
agents=""
registry_flag="--local"

help() {
echo "
Use tac_run.sh <minutes_to_start> <number_of_participants> [mode:noncontract,fetchai,ethereum] [registry:local,remote] [acn_node:local,public]

minutes_to_start - int
number_of_participants - int, min 2
mode - which mode to run Tac in: noncontract, fetchai, ethereum. noncontract is the default
registry - local or remote, local is the default
acn_node - local or public, local is the default
"
 exit 1;
}

function empty_lines {
echo "

"
}

print_options(){
start_time=$(date)
echo "Starting Tac instance:

Start time: ${start_time}

* minutes_to_start - ${minutes_to_start}
* number_of_participants - ${participants}
* mode - ${mode}
* registry - ${registry}
* acn_node - ${acn_node}
* tac_name(autogenerated) - ${tac_name}


"
}

check_options(){
	if [ -z "$minutes_to_start" ]; then
		minutes_to_start=1
	fi
	
	if [ -z "$participants" ]; then
		participants=2
	fi
	
	case "$registry" in
	  local)
        registry_flag="--local"
	    ;;
	  remote)
        registry_flag="--remote"
	    ;;
	  "")
	  	registry=local
	  	registry_flag="--local"
	    ;;
	  *)
	  	echo "Incorrect registry!"
	    help
	    ;;
	esac
	
	
	case "$acn_node" in
	  local)
	     PEER=""
	    ;;
	  public)
	     PEER="/dns4/acn.fetch.ai/tcp/9001/p2p/16Uiu2HAmVWnopQAqq4pniYLw44VRvYxBUoRHqjz1Hh2SoCyjbyRW"
	    ;;
	  "")
	  	acn_node=local
	  	PEER=""
	    ;;
	  *)
	    echo "Incorrect acn node option!"
	    help
	    ;;
	esac
	
	case "$mode" in
	  fetchai|ethereum)
	  	key_type=$mode
	  	tac_controller_name="tac_controller_contract"
	  	tac_controller_skill_name="tac_control_contract"
	  	tac_participant_name="tac_participant_contract"
	    ;;
	  ""|noncontract)
	  	key_type=fetchai
	  	mode="noncontract"
	  	tac_controller_name="tac_controller"
	  	tac_controller_skill_name="tac_control"
	  	tac_participant_name="tac_participant"
	    ;;
	  *)
	    echo "Incorrect mode option!"
	    help
	    ;;
	esac
}

generate_keys(){
	echo "Generating ${key_type} keys..."
	aea generate-key $key_type
	aea add-key $key_type ${key_type}_private_key.txt
	aea config set agent.default_ledger ${key_type}

	aea generate-key fetchai fetchai_connection_private_key.txt
	aea add-key fetchai fetchai_connection_private_key.txt --connection

	echo "Keys are generated."
}
set_p2p_connection(){
	peer=$1
	port=$2
	public_url=$3
	if [ -z "$peer" ]; then
		# no peer provided
		peer=
	else
		peer="\"$peer\""
	fi

	if [ -z "$public_url" ]; then
		# no peer provided
		public_url=null
	else
		public_url="\"$public_url\""
	fi

	json=$(printf '{"delegate_uri": null, "entry_peers": [%s], "local_uri": "127.0.0.1:1%0.4d", "public_uri": %s}' "$peer" "$port" "$public_url")
	aea config set --type dict vendor.fetchai.connections.p2p_libp2p.config "$json"
	aea config get vendor.fetchai.connections.p2p_libp2p.config
	json=$(printf '[{"identifier": "acn", "ledger_id": "%s", "not_after": "2030-01-01", "not_before": "2022-01-01", "public_key": "fetchai", "message_format": "{public_key}", "save_path": ".certs/conn_cert.txt"}]' "$key_type")
	aea config set --type list vendor.fetchai.connections.p2p_libp2p.cert_requests "$json"
}

set_tac_name(){
	echo "Set tac name"
	skill_name=$1
	json=$(printf '{"key": "tac", "value": "%s"}' $tac_name)
	aea config set --type dict vendor.fetchai.skills.$skill_name.models.parameters.args.service_data "$json"
	aea config get vendor.fetchai.skills.$skill_name.models.parameters.args.service_data
}
set_aea(){
	aea config set vendor.fetchai.connections.soef.config.chain_identifier ethereum
	aea install
	aea build
	aea issue-certificates
}

set_PEER(){
	case "$acn_node" in
	  public)
	  # do nothing
	    ;;
	  local)
	  	PEER=`aea get-multiaddress fetchai -c -i fetchai/p2p_libp2p:0.25.0 -u public_uri`
	    ;;
	esac
}
create_controller(){
	echo "Creating controller..."

	rm -rf $tac_controller_name
	aea fetch $registry_flag fetchai/$tac_controller_name:latest

	cd $tac_controller_name

	empty_lines
	generate_keys

	empty_lines
	set_p2p_connection "$PEER" "0" "127.0.0.1:10000"

	empty_lines
	set_tac_name "$tac_controller_skill_name"
	
	set_aea	
	set_PEER
	cd ..
}

create_participants(){
	for i in $(seq $participants);
	do
		create_participant $i
	done
}

create_participant(){
	i=$1
	
	agent=tac_participant_$i
	agents=$(echo $agent $agents)
	rm -rf $agent
	aea  fetch $registry_flag fetchai/${tac_participant_name}:latest --alias $agent
	cd $agent
	
	empty_lines
	generate_keys

	empty_lines
	set_p2p_connection "$PEER" "$i"

	aea config set vendor.fetchai.skills.tac_participation.models.game.args.search_query.search_value $tac_name
	aea config get vendor.fetchai.skills.tac_participation.models.game.args.search_query
		
	set_aea	
	cd ..
}

set_tac_time(){
	empty_lines
	time_diff=$(printf '+%sM' "$minutes_to_start")
	datetime_now=$(date "+%d %m %Y %H:%M")
	datetime_start=$([ "$(uname)" = Linux ] && date --date="$minutes_to_start minutes" "+%d %m %Y %H:%M" ||date -v $time_diff "+%d %m %Y %H:%M")
	echo "Now:" $datetime_now "Start:" $datetime_start
	cd $tac_controller_name
	aea config set vendor.fetchai.skills.${tac_controller_skill_name}.models.parameters.args.registration_start_time "$datetime_start"
	echo "Start time set:" $(aea config get vendor.fetchai.skills.${tac_controller_skill_name}.models.parameters.args.registration_start_time)
	cd ..
}

create_agents(){
	create_controller
	empty_lines
	create_participants
	empty_lines
	set_tac_time
	echo "agents created"
}

run_agents(){
	empty_lines
	echo "Use \"aea launch $tac_controller_name $agents\" to run agents"
}



check_options
print_options
create_agents
run_agents
