// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.7;

contract ControllerSwarm {

    struct Position {
        int32 x;
        int32 y;
    }

    struct ByzantineIndex {
        int32 nonByzIndex;
        int32 byzIndex;
    }

    struct Map {
        int[] cells;
        ByzantineIndex byzInd;
        address agent;
    }

    int32 nonByzThreshold;
    int32 byzThreshold;
    int width;
    int length;
    int32 missions;
    Map[] submittedMaps;
    mapping(address => int32) agents;

    constructor(int32 width_, int32 length_, int32 nonByzThreshold_, int32 byzThreshold_) {
        nonByzThreshold = nonByzThreshold_;
        byzThreshold = byzThreshold_;
        length = length_;
        width = width_;
        missions = 0;
    }

    function register() public {
        require(agents[msg.sender] == 0, "Agent is already registerd!");
        agents[msg.sender] = 1;
    }

    function requestMission() public returns(int32 missionID, Position memory pos) {
        require(agents[msg.sender] != 0, "Agent is not registerd!");
        require(agents[msg.sender] != 2, "Agent is marked az byzantine!");

        missions += 1;
        pos.x = 0; // some kind of randomness
        pos.y = 0; // some kind of randomness

        missionID = missions;
    }

    function abs(int x) private pure returns (int) {
        return x >= 0 ? x : -x;
    }

    function submitMap(int[] memory mapCells, int32 missionID) public {
        require(int(mapCells.length) == length * width, "Wrong input map size!");
        require(missionID > missions, "Wrong mission ID is proposed!");

        ByzantineIndex memory byzInd;
        for(uint i = 0; i < submittedMaps.length; i++) {
            Map storage map = submittedMaps[i];
            int32 byzVal = 0;
            int32 nonByzVal = 0;

            for(uint j = 0; int(j) < width*length; j++) {
                int diff = abs(mapCells[j] - map.cells[j]);
                if(diff == 0){
                    nonByzVal += 1;
                } else if(diff == 2) {
                    byzVal += 1;
                }
            }

            if(nonByzVal > nonByzThreshold) {
                map.byzInd.nonByzIndex += 1;
                byzInd.nonByzIndex += 1;
            } 
            if(byzVal > byzThreshold) {
                map.byzInd.byzIndex += 1;
                byzInd.byzIndex += 1;
            }
            
            if(map.byzInd.byzIndex > map.byzInd.nonByzIndex) {
                agents[map.agent] = 2;
            }
        }

        submittedMaps.push(Map({
            cells: mapCells,
            byzInd: byzInd,
            agent: msg.sender
        }));

        if(byzInd.byzIndex > byzInd.nonByzIndex) {
            agents[msg.sender] = 2;
        }
    } 
}
