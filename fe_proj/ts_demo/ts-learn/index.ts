// 基础类型
const msg: string = 'Hello World!!!!';
let msg1: string = '123';
const msgStr: string = `${msg} ${msg1}`;
console.log(msgStr);

// boolean
const bool: boolean = true;

// object
const obj: object = {}
const obj1: {} = {}
const obj2: { msg: string, num: number } = {msg: 'hello', num: 0}
const obj3: { msg: string, num?: number } = {msg: 'hello'}
obj3.num = 0
obj3.num = 123
console.log(obj3)

// array
const arr: [] = []
const arr3: string[] = ['123123']
const arr4: (string | number)[] = ['2233', 22222]
const arr1: Array<string> = ['123']
const arr2: Array<string | number> = ['2233', 222]

