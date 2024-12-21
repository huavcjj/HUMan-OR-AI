"use client";

import React, { useState, useEffect } from "react";
import { Button } from "./components/ui/button";
import { Input } from "./components/ui/input";
import { Textarea } from "./components/ui/textarea";

export default function Home() {
  const [keyword, setKeyword] = useState("");
  const [isMatching, setIsMatching] = useState(false);
  const [isAnswering, setIsAnswering] = useState(false);
  const [theme, setTheme] = useState("");
  const [answer, setAnswer] = useState("");
  const [timeLeft, setTimeLeft] = useState(60);
  const [isTimeUp, setIsTimeUp] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);

  const Postkeyword = async () => {
    const res = await fetch("POST http://localhost:8080/game/start", {
      method: "POST",
      headers: {
        "Content-Type": "apllication/json",
      },
      body: JSON.stringify({ passcode: keyword }),
    });
  };

  const handleStart = () => {
    if (keyword.trim() !== "") {
      setIsMatching(true);
      setTimeout(() => {
        setIsMatching(false);
        setIsLoading(true);
      }, 6000);
    }
  };

  const handleThemeSubmit = () => {
    if (theme.trim() !== "") {
      setIsMatching(true);
      setTimeout(() => {
        setIsMatching(false);
        setIsLoading(false);
        setIsAnswering(true);
        setTimeLeft(60);
        setIsTimeUp(false);
      }, 3000);
    }
  };
  const handleSubmit = () => {
    if (answer.trim() !== "" && !isTimeUp) {
      console.log("回答が提出されました:", answer);
      // ここで回答を送信するロジックを実装します
    }
  };

  useEffect(() => {
    let timer: NodeJS.Timeout;
    if (isAnswering && timeLeft > 0) {
      timer = setInterval(() => {
        setTimeLeft((prevTime) => prevTime - 1);
      }, 1000);
    } else if (timeLeft === 0) {
      setIsTimeUp(true);
    }
    return () => clearInterval(timer);
  }, [isAnswering, timeLeft]);

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
            {!isMatching && !isAnswering && !isSubmitting ? (
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
            ) : isMatching || isLoading || isSubmitting ? (
              <div className="flex flex-col items-center justify-center space-y-4">
                <div className="text-[#ffd700] text-3xl font-bold">
                  {isMatching && !isSubmitting
                    ? "マッチング中..."
                    : isLoading
                    ? "お題を取得中..."
                    : "回答を送信中..."}
                </div>
                <div className="animate-spin rounded-full h-16 w-16 border-t-4 border-[#ffd700]"></div>
                {isMatching && !isSubmitting && (
                  <div className="text-[#ffd700] text-xl">
                    あいことば: {keyword}
                  </div>
                )}
              </div>
            ) : error ? (
              <div className="flex flex-col items-center justify-center space-y-4">
                <div className="text-[#ffd700] text-3xl font-bold">エラー</div>
                <div className="text-[#ffd700] text-xl">{error}</div>
                <Button
                  onClick={fetchTheme}
                  className="bg-[#ffd700] hover:bg-[#ffec80] text-black font-bold py-2 px-4 rounded"
                >
                  再試行
                </Button>
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
                <div className="w-full">
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
