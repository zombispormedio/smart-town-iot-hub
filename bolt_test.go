package main
import (
    "testing"
    "github.com/boltdb/bolt"
)



func TestMain(m *testing.T){
    
     db, err := bolt.Open(".cache/main.db", 0600, nil)
    if err != nil {
       m.Fatal(err);
    }
    defer db.Close()
    
    db.Update(func(tx *bolt.Tx) error{
        var Error error
         b, _ :=tx.CreateBucketIfNotExists([]byte("Config"))
      
        v := string(b.Get([]byte("identifier")))
        m.Log(v)
        
        if v==""{
             m.Log("identifier changed")
            Error=b.Put([]byte("identifier"), []byte("n6NH976vNOHlWQwGH83uvXS9bTsrUtYb"))
        }
         
        return Error
    })
    
  m.Log("hello")
    
}