package usescases

import "testing"

func TestOrderInteractor_Add(t *testing.T) {
	orderId := 1
	userId := 1
	itemId := 3
	o := getFakeOrderInteractor()
	t.Log("Given the need to test add an item to order.")
	{
		t.Logf("\tWhen save an item with itemId %d and userId %d and orderId %d", itemId, userId, orderId)
		{
			err := o.Add(userId, orderId, itemId)
			if err != nil {
				t.Fatal("\t\tShould be able to add item to order.", ballotX, err)
			}
			t.Log("\t\tShould be able to add item to order.", checkMark)
		}
	}
}
