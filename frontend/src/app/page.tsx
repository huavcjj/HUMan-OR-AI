"use client";

import React, { useState, useEffect } from "react";
import { Button } from "./components/ui/button";
import { Input } from "./components/ui/input";
import { Textarea } from "./components/ui/textarea";
import { RadioGroup, RadioGroupItem } from "./components/ui/radio-group";
import { Label } from "./components/ui/label";
import corkBoard from "./images/cork-board.png";

export default function Home() {
  const [keyword, setKeyword] = useState("");
  const [keyRes, setKeyRes] = useState({ id: 0, passcode: "abc" });
  const [theme, setTheme] = useState("");
  const [userTheme, setUserTheme] = useState("");
  const [answer, setAnswer] = useState("");
  const [timeLeft, setTimeLeft] = useState(60);
  const [gameState, setGameState] = useState<
    | "input"
    | "matching"
    | "themeInput"
    | "loading"
    | "answering"
    | "submitting"
    | "judging"
    | "results"
    | "finished"
  >("input");
  const [error, setError] = useState<string | null>(null);
  const [aiAnswer, setAiAnswer] = useState("");
  const [opponentAnswer, setOpponentAnswer] = useState("");
  const [selectedAnswer, setSelectedAnswer] = useState<
    "ai" | "opponent" | null
  >(null);
  const [userResult, setUserResult] = useState({
    selectedAnswer: "",
    isCorrect: false,
  });
  const [opponentResult, setOpponentResult] = useState({
    selectedAnswer: "",
    isCorrect: false,
  });
  const [resKeyword, setResKeyWord] = useState<any>({});

  const PostKeyword = async () => {
    try {
      const timeoutPromise = new Promise((_, reject) =>
        setTimeout(() => reject(new Error("Timeout")), 20000)
      );
      const fetchPromise = await fetch("http://localhost:8080/game/start", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ passcode: keyword }),
      });

      const res = await Promise.race([fetchPromise, timeoutPromise]);
      const data = await res.json();
      setKeyRes(data);
      if (data && data.id) {
        setGameState("themeInput");
      } else {
        setError("マッチングに失敗しました。もう一度お試しください。");
        setGameState("input");
      }
    } catch (error) {
      console.error("Error:", error);
      setError("マッチングに失敗しました。もう一度お試しください。");
      setGameState("input");
    }
    console.log();
  };

  const PostTheme = async () => {
    const res = await fetch("http://localhost:8080/player/topic", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        id: keyRes?.id,
        topic: theme,
      }),
    });
    console.log(res);
  };

  const GetTopic = async () => {
    const data = await fetch(
      `http://localhost:8080/opponent/topic?id=1&passcode=${keyRes.passcode}`
    );
    const res = await data.json();
    console.log(res);
  };

  const PostAnswer = async () => {
    const res = await fetch("http://localhost:8080/player/topic/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        id: keyRes?.id,
        passcode: keyRes?.passcode,
        answer: answer,
      }),
    });
    console.log(res);
  };

  const GetGPTsAnswer = async () => {
    const data = await fetch(`http://localhost:8080/answers?id=${keyRes.id}`);
    const res = await data.json();
    console.log(res);
  };

  const handleStart = () => {
    if (keyword.trim() !== "") {
      setGameState("matching");
      setError(null);
      PostKeyword();
    }
  };

  const handleThemeSubmit = () => {
    if (userTheme.trim() !== "") {
      setGameState("loading");
      setTimeout(() => {
        setTheme(userTheme);
        setGameState("answering");
        setTimeLeft(60);
      }, 3000);
    }
  };

  const handleSubmit = () => {
    if (answer.trim() !== "" && gameState === "answering") {
      setGameState("submitting");
      setTimeout(() => {
        console.log("回答が提出されました:", answer);
        // Simulate receiving AI and opponent answers
        setAiAnswer("AIによる模範回答です。");
        setOpponentAnswer("対戦相手の面白い回答です。");
        setGameState("judging");
      }, 3000);
    }
  };

  const handleJudgment = () => {
    if (selectedAnswer) {
      console.log("選択された回答:", selectedAnswer);
      // Simulate results
      const userIsCorrect = Math.random() < 0.5;
      const opponentIsCorrect = Math.random() < 0.5;
      setUserResult({
        selectedAnswer: selectedAnswer === "ai" ? "A" : "B",
        isCorrect: userIsCorrect,
      });
      setOpponentResult({
        selectedAnswer: Math.random() < 0.5 ? "A" : "B",
        isCorrect: opponentIsCorrect,
      });
      setGameState("results");
    }
  };

  useEffect(() => {
    let timer: NodeJS.Timeout;
    if (gameState === "answering" && timeLeft > 0) {
      timer = setInterval(() => {
        setTimeLeft((prevTime) => prevTime - 1);
      }, 1000);
    } else if (timeLeft === 0 && gameState === "answering") {
      setGameState("submitting");
    }
    return () => clearInterval(timer);
  }, [gameState, timeLeft]);

  const getTimerColor = () => {
    if (timeLeft > 30) return "text-green-500";
    if (timeLeft > 10) return "text-yellow-500";
    return "text-red-500";
  };

  return (
    <div className="min-h-screen bg-[#f5efd6] flex flex-col items-center justify-center p-4 relative">
      <div
        className="w-full max-w-5xl aspect-[16/9] bg-[#e6dbb7] rounded-lg shadow-lg flex flex-col items-center justify-center relative overflow-hidden z-10 border-8 border-[#8b4513]"
        style={{
          backgroundImage: `url(${corkBoard})`,
          backgroundSize: "cover",
        }}
      >
        {/* 中央エリア */}
        <div className="w-full h-full flex flex-col items-center justify-center relative">
          {/* 背景の模様 */}

          {/* 各画面タイトル（中央エリアの内側、コンテンツエリアの外側） */}
          {gameState === "input" && (
            <div className="flex flex-col items-center space-y-4 w-full max-w-2xl m-8">
              <label
                htmlFor="keyword"
                className="text-[#143a1b] text-5xl font-bold"
              >
                あいことばでマッチング
              </label>
            </div>
          )}

          {gameState === "themeInput" && (
            <div className="flex flex-col items-center space-y-4 w-full max-w-xs m-8">
              <label
                htmlFor="theme"
                className="text-[#ffd700] text-xl font-bold"
              >
                お題を考えてください
              </label>
            </div>
          )}

          {gameState === "answering" && (
            <div className="flex flex-col items-center space-y-4 w-full max-w-xs m-8">
              <div className="text-[#ffd700] text-3xl font-bold mb-4">
                大喜利の答え
              </div>
              <div className={`text-4xl font-bold ${getTimerColor()}`}>
                残り{timeLeft}秒
              </div>
              <div className="bg-white/90 border-2 border-[#ffd700] rounded p-4 w-full">
                <h2 className="text-2xl font-bold mb-2">お題:</h2>
                <p className="text-xl">{theme}</p>
              </div>
            </div>
          )}

          {gameState === "judging" && (
            <div className="flex flex-col items-center space-y-4 w-full max-w-xs m-8">
              <div className="text-[#ffd700] text-3xl font-bold mb-4">
                人間の回答はどちら？
              </div>
              <div className="bg-white/90 border-2 border-[#ffd700] rounded p-4 w-full mb-4">
                <h2 className="text-2xl font-bold mb-2">お題:</h2>
                <p className="text-xl">{theme}</p>
              </div>
            </div>
          )}
          {gameState === "results" && (
            <div className="flex flex-col items-center space-y-4 w-full max-w-xs m-8">
              <div className="text-[#ffd700] text-3xl font-bold mb-4">
                結果発表
              </div>

              <Button
                onClick={() => setGameState("finished")}
                className="mt-4 bg-[#ffd700] hover:bg-[#ffec80] text-black font-bold py-2 px-4 rounded"
              >
                ゲーム終了
              </Button>
            </div>
          )}
          {gameState === "finished" && (
            <div className="flex flex-col items-center space-y-4 w-full max-w-xs m-8">
              <div className="text-[#ffd700] text-3xl font-bold">
                ゲーム終了！
              </div>
            </div>
          )}
          {/* コンテンツエリア */}
          <div className="bg-white border-8 border-blue-400 rounded-3xl p-24 space-y-4 relative z-10 shadow-md w-4/6 max-w-lg h-[40vh]">
            {/* コンテンツ */}
            <div className="text-black text-center">
              {gameState === "input" && (
                <div className="flex flex-col items-center space-y-8 w-full max-w-xs">
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
                    className="bg-[#ffd700] hover:bg-[#ffec80] text-black font-bold py-6 px-6 rounded"
                  >
                    スタート
                  </Button>
                </div>
              )}
              {gameState === "matching" && (
                <div className="flex flex-col items-center justify-center space-y-4">
                  <div className="text-[#ffd700] text-3xl font-bold">
                    マッチング中...
                  </div>
                  <div className="animate-spin rounded-full h-16 w-16 border-t-4 border-[#ffd700]"></div>
                  <div className="text-[#ffd700] text-xl">
                    あいことば: {keyword}
                  </div>
                </div>
              )}
              {gameState === "themeInput" && (
                <div className="flex flex-col items-center space-y-4 w-full max-w-xs">
                  <Input
                    id="theme"
                    type="text"
                    value={userTheme}
                    onChange={(e) => setUserTheme(e.target.value)}
                    className="bg-white/90 border-2 border-[#ffd700] text-black text-center"
                    placeholder="お題を入力"
                  />
                  <Button
                    onClick={handleThemeSubmit}
                    disabled={userTheme.trim() === ""}
                    className="bg-[#ffd700] hover:bg-[#ffec80] text-black font-bold py-2 px-4 rounded"
                  >
                    お題を送信
                  </Button>
                </div>
              )}
              {gameState === "loading" && (
                <div className="flex flex-col items-center justify-center space-y-4">
                  <div className="text-[#ffd700] text-3xl font-bold">
                    お題を取得中...
                  </div>
                  <div className="animate-spin rounded-full h-16 w-16 border-t-4 border-[#ffd700]"></div>
                </div>
              )}
              {gameState === "answering" && (
                <div className="flex flex-col items-center justify-center space-y-4 w-full">
                  <div className="w-full">
                    <h3 className="text-[#ffd700] text-xl font-bold mb-2">
                      あなたの回答
                    </h3>
                    <Textarea
                      value={answer}
                      onChange={(e) => setAnswer(e.target.value)}
                      placeholder="回答を入力してください"
                      className="w-full h-32 bg-white/90 border-2 border-[#ffd700] text-black p-2 rounded"
                    />
                    <Button
                      onClick={handleSubmit}
                      disabled={answer.trim() === ""}
                      className="mt-2 bg-[#ffd700] hover:bg-[#ffec80] text-black font-bold py-2 px-4 rounded"
                    >
                      回答を送信
                    </Button>
                  </div>
                </div>
              )}
              {gameState === "submitting" && (
                <div className="flex flex-col items-center justify-center space-y-4">
                  <div className="text-[#ffd700] text-3xl font-bold">
                    回答を送信中...
                  </div>
                  <div className="animate-spin rounded-full h-16 w-16 border-t-4 border-[#ffd700]"></div>
                </div>
              )}
              {gameState === "judging" && (
                <div className="flex flex-col items-center justify-center space-y-4 w-full">
                  <RadioGroup
                    value={selectedAnswer || ""}
                    onValueChange={(value) =>
                      setSelectedAnswer(value as "ai" | "opponent")
                    }
                    className="w-full space-y-2"
                  >
                    <div className="flex items-center space-x-2 bg-white/90 p-4 rounded">
                      <RadioGroupItem value="ai" id="ai" />
                      <Label htmlFor="ai" className="text-black">
                        回答A: {aiAnswer}
                      </Label>
                    </div>
                    <div className="flex items-center space-x-2 bg-white/90 p-4 rounded">
                      <RadioGroupItem value="opponent" id="opponent" />
                      <Label htmlFor="opponent" className="text-black">
                        回答B: {opponentAnswer}
                      </Label>
                    </div>
                  </RadioGroup>
                  <Button
                    onClick={handleJudgment}
                    disabled={!selectedAnswer}
                    className="mt-4 bg-[#ffd700] hover:bg-[#ffec80] text-black font-bold py-2 px-4 rounded"
                  >
                    判定
                  </Button>
                </div>
              )}
              {gameState === "results" && (
                <div className="flex flex-col items-center justify-center space-y-4 w-full">
                  <div className="flex w-full justify-between">
                    <div className="bg-white/90 border-2 border-[#ffd700] rounded p-4 w-[48%]">
                      <h3 className="text-xl font-bold mb-2">あなたの結果</h3>
                      <p>選んだ回答: {userResult.selectedAnswer}</p>
                      <p>
                        判定:{" "}
                        {userResult.isCorrect ? (
                          <span className="text-green-600 font-bold">正解</span>
                        ) : (
                          <span className="text-red-600 font-bold">不正解</span>
                        )}
                      </p>
                    </div>
                    <div className="bg-white/90 border-2 border-[#ffd700] rounded p-4 w-[48%]">
                      <h3 className="text-xl font-bold mb-2">相手の結果</h3>
                      <p>選んだ回答: {opponentResult.selectedAnswer}</p>
                      <p>
                        判定:{" "}
                        {opponentResult.isCorrect ? (
                          <span className="text-green-600 font-bold">正解</span>
                        ) : (
                          <span className="text-red-600 font-bold">不正解</span>
                        )}
                      </p>
                    </div>
                  </div>
                </div>
              )}
              {gameState === "finished" && (
                <div className="flex flex-col items-center justify-center space-y-4">
                  <div className="text-[#ffd700] text-xl">お疲れ様でした。</div>
                  <Button
                    onClick={() => {
                      setGameState("input");
                      setKeyword("");
                      setTheme("");
                      setUserTheme("");
                      setAnswer("");
                      setSelectedAnswer(null);
                      setError(null);
                    }}
                    className="mt-4 bg-[#ffd700] hover:bg-[#ffec80] text-black font-bold py-2 px-4 rounded"
                  >
                    新しいゲームを始める
                  </Button>
                </div>
              )}
            </div>
          </div>
          {/* ホワイトボードマーカー */}
          <div className="absolute bottom-4 right-4 w-16 h-4 bg-black rounded-lg shadow-md flex items-center justify-center">
            <div className="w-2 h-3 bg-white rounded-sm"></div>
          </div>
          {error && <div className="text-red-500 font-bold mt-4">{error}</div>}
        </div>
      </div>
    </div>
  );
}
