package usescases

import "testing"
func TestOrderInteractor_Items(t *testing.T) {
	orderId := 1
	userId := 1
	o := getFakeOrderInteractor()
	t.Log("Given the need to test retrieve items.")
	{
		t.Logf("\tWhen get items for a userId %d and orderId %d", userId, orderId)
		{
			items, err := o.Items(userId, orderId)
			if err != nil {
				t.Fatal("\t\tShould be able to retrieve all items.", ballotX, err)
			}

			t.Log("\t\tShould be able to retrieve all items.", checkMark)

			if len(items) == 2 {
				t.Logf("\t\tShould receive two items %v", checkMark)
			} else {
				t.Errorf("\t\tShould receive two items but found %d. %v", len(items), ballotX)
			}
		}
	}
}
