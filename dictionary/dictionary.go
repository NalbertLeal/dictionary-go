package dictionary

import (
  "fmt"
  "errors"
)

type Capsule struct {
  key           interface{}
  data          interface{}
  timesAccessed int64
}

type Dictionary struct {
  size          int64
  capacity      int64
  alwaysSorted  bool
  compare       func(interface{}, interface{}) int8
  array         []*Capsule
}

func NewDictionary(sorted bool, comp func(interface{}, interface{}) int8) *Dictionary {
  var temp *Dictionary
  temp = &Dictionary{
    size: 0,
    capacity: 1000,
    alwaysSorted: sorted,
    compare: comp,
    array: make([]*Capsule, 1000),
  }

  for i := 0; i < 1000; i++ {
    temp.array[i] = nil
  }

  return temp
}

func CopyDictionary(dict *Dictionary) *Dictionary {
  var temp *Dictionary
  temp = &Dictionary{
    size: dict.GetSize(),
    capacity: dict.GetCapacity(),
    alwaysSorted: dict.GetAlwaysSorted(),
    compare: dict.GetCompare(),
    array: make([]*Capsule, dict.GetCapacity()),
  }

  var index int64
  for index = 0; index < temp.size; index++ {
    temp.array[index] = dict.GetIndex(index)
  } 

  return temp
}

func (this *Dictionary) GetSize() int64 {
  return this.size
}

func (this *Dictionary) GetCapacity() int64 {
  return this.capacity
}

func (this *Dictionary) GetAlwaysSorted() bool {
  return this.alwaysSorted
}

func (this *Dictionary) GetCompare() (func(interface{}, interface{}) int8) {
  return this.compare
}

func (this *Dictionary) GetIndex(index int64) (*Capsule) {
  if this.size > index && this.array[index] != nil {
    return this.array[index]
  } else {
    return nil
  }
}

func (this *Dictionary) Search(tempKey interface{}) (interface{}, error) {
  var cap *Capsule = nil
  var err error

  cap, _, err = this.search(tempKey)
  if err != nil {
    return cap, errors.New("This Key doesn't exist into the dictionary")
  }

  return cap.data, nil
}

func (this *Dictionary) Insert(tempKey interface{}, tempData interface{}) {
  var index int64

  _, index, _ = this.search(tempKey)

  if index != -1 {
    // update value
    this.array[index].data = tempData
  } else {
    // push Capsule with key and data into the array
    if this.size == this.capacity-1 {
      this.doubleArraySize()
    }

    this.array[this.size] = &Capsule{
      key: tempKey,
      data: tempData,
      timesAccessed: 0,
    }
    this.size++

    if this.alwaysSorted {
      this.sort()
    }

  }
}

func (this *Dictionary) Remove(tempKey interface{}) error {
  var index int64
  var err error

  _, index, err = this.search(tempKey)
  if err != nil {
    return errors.New("No Key like the required")
  }

  this.array[index] = nil

  return nil
}

func (this *Dictionary) Exist(tempKey interface{}) bool {
  var index int64
  var err error

  _, index, err = this.search(tempKey)
  if err != nil && index != -1 {
    return true
  } else {
    return false
  }
}

func (this *Dictionary) search(tempKey interface{}) (*Capsule, int64, error){
  var begin int64 = 0
  var end int64 = this.size
  var middle int64

  for begin <= end {
    middle = (begin + ((end-begin)/2))

    if this.array[middle] == nil {
      return &Capsule{}, -1, errors.New("Key not founded.")
    }

    switch this.compare(this.array[middle].key, tempKey) {
      case 0:
       return this.array[middle], middle, nil
      case 1:
        // this.array[middle].key < tempKey
        begin = middle + 1
        break
      default:
        // this.array[middle].key > tempKey
        end = middle - 1
    }
  }

  return &Capsule{}, -1, errors.New("Key not founded.")
}

func (this *Dictionary) sort() {

  var tempCapsule *Capsule
  var index int64
  var j int64

  for index = 0; index < this.size; index++ {
    tempCapsule = this.array[index]
    for j = index; j < this.size; j++ {
      if (this.array[j] == nil) || (this.array[j+1] == nil) {
        break
      }
      if (this.compare(this.array[j].key, this.array[j+1].key) == 1) {
        break
      }
      this.array[j] = this.array[j+1]
    }
    this.array[j] = tempCapsule
  }

}

func (this *Dictionary) doubleArraySize() {
  var tempArray []*Capsule = this.array

  this.array = make([]*Capsule, this.capacity*2)

  for index, value := range tempArray {
    this.array[index] = value
  }
}

func (this *Dictionary) PrintDictHead() {
  var index int8
  for index = 0; index < 5; index++ {
    fmt.Println(this.array[index])
  }
}