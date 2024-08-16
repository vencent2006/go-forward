export function success(data, message){
  return {
    code: 0,
    data,
    message
  }
}

export function error(message){
  return {
    code: -1,
    message
  }
}
