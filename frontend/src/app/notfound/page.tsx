import { useRouter } from 'next/router'
export default function page(){
    const router = useRouter();
    return(
        <div>
            <div>相手が見つかりませんでした。</div>
            <button onClick={()=>router.push("/")}>最初に戻る</button>
        </div>
    )
}