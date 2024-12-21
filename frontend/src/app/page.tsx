'use client'

import React, { useState, useEffect } from "react";
import { Button } from "./components/ui/button";
import { Input } from "./components/ui/input";
import { Textarea } from "./components/ui/textarea";

export default function Home() {
  const [keyword, setKeyword] = useState("");
  const [isMatching, setIsMatching] = useState(false);
  const [isMatched, setIsMatched] = useState(false);
  const [theme, setTheme] = useState("");
  const [answer, setAnswer] = useState("");
  const [gptAnswer, setGptAnswer] = useState("");
  const [isGptThinking, setIsGptThinking] = useState(false);
  const [timeLeft, setTimeLeft] = useState(60);
  const [isTimeUp, setIsTimeUp] = useState(false);
  const [response,setResponse] = useState({
    "id":0,
    "passcode":"s"
  })

  const Postkeyword = async()=>{
    const res = await fetch('http://localhost:8080/game/start',{
      method:'POST',
      headers:{
        'Content-Type':'apllication/json',
      },
      body:JSON.stringify({"passcode":keyword})
    });
    const responses = res.json();
    console.log(responses);
  }

  const PostTheme = async()=>{
    const res = await fetch('http://localhost:8080/game/start',{
      method:'POST',
      headers:{
        'Content-Type':'apllication/json',
      },
      body:JSON.stringify({"id":response.id})
    });
  } 
 
  const handleStart = () => {
    if (keyword.trim() !== "") {
      Postkeyword();
      setIsMatching(true);
      setTimeout(() => {
        Postkeyword();
        setIsMatching(false);
        setIsMatched(true);
        setTheme(
          "笑点でよく見る「座布団何枚」を現代風にアレンジするとどうなる？"
        );
        setIsGptThinking(true);
        setTimeLeft(60);
        setIsTimeUp(false);
      }, 6000);
    }
  };

  const handleSubmit = () => {
    if (answer.trim() !== "" && !isTimeUp) {
      console.log("回答が提出されました:", answer);
      // ここで回答を送信するロジックを実装します
    }
  };

  useEffect(() => {
    if (isGptThinking) {
      const thinkingTime = Math.random() * 5000 + 5000; // 5-10秒のランダムな時間
      const timer = setTimeout(() => {
        setGptAnswer(
          "「いいね何件！」現代のSNS時代にマッチした新しい褒め方ですね。"
        );
        setIsGptThinking(false);
      }, thinkingTime);
      return () => clearTimeout(timer);
    }
  }, [isGptThinking]);

  useEffect(() => {
    let timer: NodeJS.Timeout;
    if (isMatched && timeLeft > 0) {
      timer = setInterval(() => {
        setTimeLeft((prevTime) => prevTime - 1);
      }, 1000);
    } else if (timeLeft === 0) {
      setIsTimeUp(true);
    }
    return () => clearInterval(timer);
  }, [isMatched, timeLeft]);

  const getTimerColor = () => {
    if (timeLeft > 30) return "text-green-500";
    if (timeLeft > 10) return "text-yellow-500";
    return "text-red-500";
  };

  return (
    <div className="min-h-screen bg-[#f5efd6] flex flex-col items-center justify-center p-4 relative">
      {/* 提灯（ちょうちん）の列 */}
      <div className="absolute top-0 left-0 right-0 flex justify-center overflow-hidden">
        <div className="flex space-x-4 py-4">
          {[...Array(10)].map((_, i) => (
            <div
              key={i}
              className="w-12 h-16 bg-red-600 rounded-full flex items-center justify-center relative"
              style={{
                boxShadow: "inset 0 0 10px rgba(0,0,0,0.2)",
              }}
            >
              <div className="absolute top-0 w-8 h-2 bg-[#8b4513]"></div>
            </div>
          ))}
        </div>
      </div>

      <div className="w-full max-w-4xl aspect-[16/9] bg-[#e6dbb7] rounded-lg shadow-lg flex flex-col items-center justify-center relative overflow-hidden z-10 border-8 border-[#8b4513]">
        {/* 中央エリア */}
        <div className="w-full h-full flex flex-col items-center justify-center relative">
          {/* 背景の模様 */}
          <div
            className="absolute inset-0 opacity-20"
            style={{
              backgroundImage: `url("data:image/svg+xml,%3Csvg width='120' height='120' viewBox='0 0 120 120' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M0 0h60v60H0z' fill='%234a69bd'/%3E%3Cpath d='M60 0h60v60H60zM0 60h60v60H0z' fill='%234a69bd' fill-opacity='.5'/%3E%3Cpath d='M60 60h60v60H60z' fill='%234a69bd'/%3E%3C/svg%3E")`,
              backgroundSize: "120px 120px",
            }}
          ></div>

          {/* コンテンツエリア */}
          <div className="bg-[#cc0000] border-4 border-[#ffd700] rounded-lg p-8 space-y-4 relative z-10 w-5/6 max-w-2xl">
            {!isMatching && !isMatched ? (
              <>
                <div className="flex flex-col items-center space-y-4 w-full max-w-xs">
                  <label
                    htmlFor="keyword"
                    className="text-[#ffd700] text-xl font-bold"
                  >
                    あいことば
                  </label>
                  <Input
                    id="keyword"
                    type="text"
                    value={keyword}
                    onChange={(e) => setKeyword(e.target.value)}
                    className="bg-white/90 border-2 border-[#ffd700] text-black text-center"
                    placeholder="あいことばを入力"
                  />
                  <Button
                    onClick={handleStart}
                    disabled={keyword.trim() === ""}
                    className="bg-[#ffd700] hover:bg-[#ffec80] text-black font-bold py-2 px-4 rounded"
                  >
                    スタート
                  </Button>
                </div>
              </>
            ) : isMatching ? (
              <div className="flex flex-col items-center justify-center space-y-4">
                <div className="text-[#ffd700] text-3xl font-bold">
                  マッチング中...
                </div>
                <div className="animate-spin rounded-full h-16 w-16 border-t-4 border-[#ffd700]"></div>
                <div className="text-[#ffd700] text-xl">
                  あいことば: {keyword}
                </div>
              </div>
            ) : (
              <div className="flex flex-col items-center justify-center space-y-4 w-full">
                <div className="text-[#ffd700] text-3xl font-bold mb-4">
                  大喜利
                </div>
                <div className={`text-4xl font-bold ${getTimerColor()}`}>
                  {timeLeft}秒
                </div>
                <div className="bg-white/90 border-2 border-[#ffd700] rounded p-4 w-full">
                  <h2 className="text-2xl font-bold mb-2">お題:</h2>
                  <p className="text-xl">{theme}</p>
                </div>
                <div className="flex w-full space-x-4">
                  <div className="flex-1">
                    <h3 className="text-[#ffd700] text-xl font-bold mb-2">
                      あなたの回答
                    </h3>
                    <Textarea
                      value={answer}
                      onChange={(e) => setAnswer(e.target.value)}
                      placeholder="回答を入力してください"
                      className="w-full h-32 bg-white/90 border-2 border-[#ffd700] text-black p-2 rounded"
                      disabled={isTimeUp}
                    />
                    <Button
                      onClick={handleSubmit}
                      disabled={answer.trim() === "" || isTimeUp}
                      className="mt-2 bg-[#ffd700] hover:bg-[#ffec80] text-black font-bold py-2 px-4 rounded"
                    >
                      回答を送信
                    </Button>
                  </div>
                  <div className="flex-1">
                    <h3 className="text-[#ffd700] text-xl font-bold mb-2">
                      ChatGPTの回答
                    </h3>
                    <div className="w-full h-32 bg-white/90 border-2 border-[#ffd700] text-black p-2 rounded overflow-y-auto">
                      {isGptThinking ? (
                        <div className="flex items-center justify-center h-full">
                          <div className="animate-pulse text-lg">考え中...</div>
                        </div>
                      ) : (
                        <p>{gptAnswer}</p>
                      )}
                    </div>
                  </div>
                </div>
                {isTimeUp && (
                  <div className="text-[#ffd700] text-2xl font-bold mt-4">
                    時間切れ！
                  </div>
                )}
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}

