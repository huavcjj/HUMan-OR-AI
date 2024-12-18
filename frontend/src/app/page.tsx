import Signin from "./components/Signin";
import {useAuthState} from "react-firebase-hooks/auth"
import firebase from "firebase/compat/app"
import {auth} from "../../src/app/Firebase"

export default function Home() {
  return (
    <div className="min-h-[100vh]">
      <Signin/>
    </div>
  );
}
