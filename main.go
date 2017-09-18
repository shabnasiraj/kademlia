package main

import (
	"fmt"
	"strconv"
)

func main() {
	kademlia := NewKademlia()
	IDRTList := map[KademliaID]RoutingTable{}
	firstNode := NewContact(NewRandomKademliaID(), "localhost:8000")
	firstNodeRT := NewRoutingTable(firstNode)

	for i := 0; i < 10; i++ {
		port := 6000 + i
		a := "localhost," + strconv.Itoa(port)
		ID := NewRandomKademliaID()
		rt := NewRoutingTable(NewContact(ID, a))
		IDRTList[*ID] = *rt
	}
	for k, v := range IDRTList {
		v.AddContact(firstNode)
		closestNodes := kademlia.LookupContact(IDRTList[k].me, *firstNodeRT)
		for i := range closestNodes {
			v.AddContact(closestNodes[i])
		}
	}
	for i := range firstNodeRT.buckets {
		contactList := firstNodeRT.buckets[i]
		fmt.Println("Bucket: " + strconv.Itoa(i))
		for elt := contactList.list.Front(); elt != nil; elt = elt.Next() {
			contact := elt.Value.(Contact)
			fmt.Println(contact.String())
		}

}

}
