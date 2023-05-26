package utils

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// WriteLedger 将一个对象写入 Hyperledger Fabric 区块链账本中
// obj：要写入账本的对象，可以是任意类型的数据。
// stub：智能合约与区块链网络进行交互的接口，类型为 shim.ChaincodeStubInterface。
// objectType：对象类型，用于创建复合主键时区分不同类型的对象。
// keys：用于创建复合主键的键值列表。
// 该函数可以在智能合约中的任何位置调用，用于将任意类型的数据写入区块链账本中
func WriteLedger(obj interface{}, stub shim.ChaincodeStubInterface, objectType string, keys []string) error {
	// 使用 stub.CreateCompositeKey 方法根据 objectType 和 keys 创建一个复合主键
	var key string
	if val, err := stub.CreateCompositeKey(objectType, keys); err != nil {
		return errors.New(fmt.Sprintf("%s-创建复合主键出错 %s", objectType, err))
	} else {
		key = val
	}
	bytes, err := json.Marshal(obj) // 将 obj 序列化成 JSON 字符串
	if err != nil {
		return errors.New(fmt.Sprintf("%s-序列化json数据失败出错: %s", objectType, err))
	}
	// 调用 stub.PutState 方法将序列化后的数据写入区块链账本中
	if err := stub.PutState(key, bytes); err != nil {
		return errors.New(fmt.Sprintf("%s-写入区块链账本出错: %s", objectType, err))
	}
	return nil
}

// DelLedger 用于从 Hyperledger Fabric 区块链账本中删除一个对象
// stub：智能合约与区块链网络进行交互的接口，类型为 shim.ChaincodeStubInterface。
// objectType：对象类型，用于创建复合主键。
// keys：用于创建复合主键的键值列表
func DelLedger(stub shim.ChaincodeStubInterface, objectType string, keys []string) error {
	// 创建复合主键
	var key string
	if val, err := stub.CreateCompositeKey(objectType, keys); err != nil {
		return errors.New(fmt.Sprintf("%s-创建复合主键出错 %s", objectType, err))
	} else {
		key = val
	}
	// 使用 stub.DelState 方法将该主键对应的对象从区块链账本中删除
	if err := stub.DelState(key); err != nil {
		return errors.New(fmt.Sprintf("%s-删除区块链账本出错: %s", objectType, err))
	}
	return nil
}

// GetStateByPartialCompositeKeys 根据复合主键查询数据(适合获取全部，多个，单个数据)
// stub：智能合约与区块链网络进行交互的接口，类型为 shim.ChaincodeStubInterface。
// objectType：对象类型，用于创建复合主键。
// keys：用于创建复合主键的键值列表，一个字符串数组
// 该函数返回一个二维字节数组，表示从账本中获取到的数据
func GetStateByPartialCompositeKeys(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results [][]byte, err error) {
	if len(keys) == 0 {
		// 传入的keys长度为0，则查找并返回所有数据
		resultIterator, err := stub.GetStateByPartialCompositeKey(objectType, keys)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s-获取全部数据出错: %s", objectType, err))
		}
		defer resultIterator.Close()

		//检查返回的数据是否为空，不为空则遍历数据，否则返回空数组
		for resultIterator.HasNext() {
			val, err := resultIterator.Next()
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-返回的数据出错: %s", objectType, err))
			}

			results = append(results, val.GetValue())
		}
	} else {
		// 传入的keys长度不为0，查找相应的数据并返回
		for _, v := range keys {
			// 创建组合键
			key, err := stub.CreateCompositeKey(objectType, []string{v})
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-创建组合键出错: %s", objectType, err))
			}
			// 从账本中获取数据stub.GetState(key)
			bytes, err := stub.GetState(key) // 该方法会根据给定的主键从区块链账本中获取相应的数据，并返回一个字节数组类型的结果
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-获取数据出错: %s", objectType, err))
			}

			if bytes != nil {
				results = append(results, bytes)
			}
		}
	}

	return results, nil
}

// GetStateByPartialCompositeKeys2 根据复合主键查询数据(适合获取全部或指定的数据)
// 与 GetStateByPartialCompositeKeys 不同的是，该函数不会根据 keys 的长度来区分是获取全部数据还是获取指定数据。
// 它仅仅是从账本中获取符合条件的数据，并将其返回。
// 因此，该函数通常用于在 Hyperledger Fabric 智能合约中查询符合特定条件的复合主键对应的数据，可以根据不同的 keys 参数返回不同的结果。
func GetStateByPartialCompositeKeys2(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results [][]byte, err error) {
	resultIterator, err := stub.GetStateByPartialCompositeKey(objectType, keys)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s-获取全部数据出错: %s", objectType, err))
	}
	defer resultIterator.Close()

	//检查返回的数据是否为空，不为空则遍历数据，否则返回空数组
	for resultIterator.HasNext() {
		val, err := resultIterator.Next()
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s-返回的数据出错: %s", objectType, err))
		}

		results = append(results, val.GetValue())
	}
	return results, nil
}
