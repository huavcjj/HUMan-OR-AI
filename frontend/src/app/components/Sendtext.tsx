import React, { ChangeEvent, FormEvent, useState } from 'react'
import {db,auth} from "../Firebase"
import firebase from "firebase/compat/app"

interface info{
    uid?:string
}

const Sendtext = () => {
    const [messages,setMessages] = useState<string>("");
    const sendMessage = (e:FormEvent)=>{
        e.preventDefault();
        if(!messages)return;
        const {uid,photoURL} = auth.currentUser;
        db.collection("messages").add({
            text:messages,
            uid,
            photoURL,
            createdAt: firebase.firestore.FieldValue.serverTimestamp(),
        })
        setMessages("");
    }
  return (
    <div>
        <form onSubmit={sendMessage}>
            <input 
            className="border border-blue-400 outline-none" 
            type="text"
            onChange={(e:ChangeEvent<HTMLInputElement>)=>setMessages(e.target.value)}
            value={messages}
            />
        </form>
    </div>
  )
}

export default Sendtext