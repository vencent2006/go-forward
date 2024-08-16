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

export function wrapperResponse(p, msg){
  return p
    .then(data=>success(data,msg))
    .catch(err=>error(err.message))
}
