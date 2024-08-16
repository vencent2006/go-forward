export function success(data, msg){
  return {
    code: 0,
    data,
    msg
  }
}

export function error(msg){
  return {
    code: -1,
    msg
  }
}
