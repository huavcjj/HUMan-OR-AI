import React from 'react'

const dataFetch = async ()=>{
  const api = await fetch('http://localhost:8080');
  const res = await api.json();
  return res;
}

const Test = async () => {
  const data = await dataFetch();
  console.log(data)
  return (
    <div>Test</div>
  )
}

export default Test