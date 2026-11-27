package json

import (
	"encoding/json"
	"os"
)

func TryEncode() error {
    data := map[string]interface{}{
        "1": "one",
        "2": "two",
        "3": "three",
    }
    err := json.NewEncoder(os.Stdout).Encode(&data)
    if err != nil {
        return err
    }

    return nil
}

func TryDecode(path string) error {
    myFile, err := os.Open(path)
    if err != nil {
        return err
    }
    defer myFile.Close()

    var result map[string]interface{}
    json.NewDecoder(myFile).Decode(&result)

    return nil
}