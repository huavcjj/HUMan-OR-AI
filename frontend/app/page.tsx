'use client'
import { ChangeEvent, useState } from "react";
import Test from "../components/Test";
import io from "socket.io-client"

const socket = io("http://localhost:5000");

export default function Home() {
  const [message,setMessage] = useState<string>("");
  const handleSend = ()=>{
    socket.emit("send_message",{message:message});
    setMessage("");
  }
  return (
    <div>
      <div className="">
        <h2 className="font-bold">通信テスト</h2>
        <input type="text" className="bg-blue-200 rounded-md"  onChange={(e:ChangeEvent<HTMLInputElement>)=>setMessage(e.target.value)} value={message}/>
        <button onClick={handleSend} className="rounded-md bg-black text-white w-[60px] h-[40px] text-center hover:scale-110 transition-all">送信</button>
      </div>
      <Test/>
    </div>
  );
}
