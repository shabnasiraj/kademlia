package main

import (
	"fmt"
	"strconv"
)


//type Kademlia struct {
 // routes *RoutingTable;
 // NetworkId string;
//}

func main() {
func NewKademlia(self *Contact, networkId string) (ret *Kademlia) {
  ret = new(Kademlia);
  ret.routes = NewRoutingTable(self);
  ret.NetworkId = networkId;
  return;


//nodes

for i := 0; i < 10; i++ {
		port := 8001 + i
		a := "localhost" + strconv.Itoa(port)
		ID := NewRandomKademliaID()
		//ID := NewKademliaID(nodeIDs[i])
		rt := NewRoutingTable(NewContact(ID, a))
		IDRTList[*ID] = *rt
}
