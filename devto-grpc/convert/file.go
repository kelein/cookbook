package convert

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// WriteBinaryFile writes protobuf into binary file
func WriteBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("marshal proto error: %w", err)
	}
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("write file error: %w", err)
	}
	return nil
}

// ReadBinaryFile reads proto message from binary file
func ReadBinaryFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read file error: %w", err)
	}
	if err := proto.Unmarshal(data, message); err != nil {
		return fmt.Errorf("unmarshal proto error: %w", err)
	}
	return nil
}

// MarshalJSON convert proto message into JSON byte
func MarshalJSON(message proto.Message) ([]byte, error) {
	m := protojson.MarshalOptions{
		Indent:            " ",
		EmitDefaultValues: true,
		UseEnumNumbers:    false,
	}
	return m.Marshal(message)
}

// WriteJSONFile write proto message into JSON file
func WriteJSONFile(message proto.Message, filename string) error {
	data, err := MarshalJSON(message)
	if err != nil {
		return fmt.Errorf("marshal proto to json error: %w", err)
	}
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("write JSON byte error: %w", err)
	}
	return nil
}
