'use client'

import React, { useEffect, useState } from 'react'
import {useAuthState} from "react-firebase-hooks/auth"
import firebase from "firebase/compat/app"
import {auth,db} from "../Firebase"
import Auth from "./Auth"
import { getDocs } from 'firebase/firestore'

const Signin = () => {
  const [user] = useAuthState(auth as any);
  function SignInWithGoogle(){
    const provider = new firebase.auth.GoogleAuthProvider();
    provider.setCustomParameters({
      prompt: "select_account",
    });
    auth.setPersistence(firebase.auth.Auth.Persistence.SESSION)
    .then(()=>{
      return auth.signInWithPopup(provider);
    })
  }
  function SignInWIthGithub(){
    const provider = new firebase.auth.GithubAuthProvider();
    provider.setCustomParameters({
      prompt:"select_account",
    })
    auth.signInWithRedirect(provider);
  }

  return (
  <div className='w-full  '>
    {user?    
        <button className='w-[70px] h-[50px] bg-black text-white' onClick={()=>auth.signOut()}>
          ログアウト
        </button> :
        <button className='w-[70px] h-[50px] bg-black text-white' onClick={()=>SignInWithGoogle()}>
          googleでログイン
        </button>
  }
  {/* {user&&<Auth/>} */}
  </div>
  )
}

export default Signin