'use client'

import React, { useEffect, useState } from 'react'
import {auth} from "../Firebase"
import {db} from "../Firebase"
import {deleteDoc,getDocs} from 'firebase/firestore';
import firebase from "firebase/compat/app"
import Sendtext from './Sendtext'
import Image from 'next/image'
import { type ClassValue, clsx } from "clsx"

type Message = {
  text?: string
  photoURL?:string
  uid?:string
  createdAt: firebase.firestore.Timestamp;
};

const Auth = () => {
  const [messages,setMessages] = useState<Message[]>([]);
  const deleteDocument = async () => {
    const snapshot = await getDocs(db.collection("messages"));
    const deletePromises = snapshot.docs.map((doc) => deleteDoc(doc.ref)); 
    await Promise.all(deletePromises); 
  };
  useEffect(()=>{
    deleteDocument().then(()=>{
      db.collection("messages").
      orderBy("createdAt").
      limit(50).
      onSnapshot((snapshot)=>setMessages(snapshot.docs.map(doc => doc.data() as Message)))
    })
  },[])
  console.log(messages)
  return (
      <div className='w-full min-h-[100vh] flex items-start flex-col'>
      <Sendtext/>
        {messages?.map((doc,value)=>{
          return(
            <div className={'flex items-center mb-6'} key={value}>
              {doc.photoURL&&
                <Image
                  src={doc?.photoURL}
                  alt="ICON"
                  width={40}
                  height={40}
                  style={
                    {borderRadius:"50%",
                     marginRight:"10px"  
                  }}
                 />
                 }
              <div className={clsx(
                'inline-block min-w-[60px]',
                doc.uid===auth.currentUser?.uid?'bg-blue-300':'bg-blue-800'
                )}>{doc.text}
              </div>
            </div>
          )
        })}
      </div>
  )
}

export default Auth