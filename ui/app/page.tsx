"use client";

import { useEffect, useState } from "react";

type Finding = {
  control: string;
  status: string;
  severity: string;
  evidence: string;
  remediation: string;
  checkedAt: string;
  resource: string;
};

const sevColor: Record<string, string> = {
  critical: "bg-red-600",
  high: "bg-orange-500",
  medium: "bg-yellow-500",
  low: "bg-blue-500",
};

export default function Home() {
  const [findings, setFindings] = useState<Finding[] | null>(null);
  const [error, setError] = useState<string | null>(null);

  const load = (signal?: AbortSignal) => {
    fetch("/api/findings", { signal })
      .then((r) => {
        if (!r.ok) throw new Error(`HTTP ${r.status}`);
        return r.json();
      })
      .then(setFindings)
      .catch((e) => setError(e.message));
  };

  useEffect(() => { 
    const controller = new AbortController();
    load(controller.signal);

    return () => {
      controller?.abort();
    };
  }, []);

  if (error) return <main className="p-8 text-red-400">Failed: {error}</main>;
  if (!findings) return <main className="p-8">Scanning AWS...</main>;

  return (
    <main className="min-h-screen bg-zinc-950 p-8 text-zinc-100">
      <div className="mb-6 flex items-center justify-between">
  <h1 className="text-2xl font-bold">CloudProof</h1>
  <button onClick={() => {
    setFindings(null);
    setError(null);
    load();
  }} className="rounded bg-zinc-800 px-3 py-1 text-sm hover:bg-zinc-700">
    Re-scan
  </button>
  </div>
      <div className="grid gap-4 md:grid-cols-2">
        {findings.map((f) => (
          <div key={f.control} className="rounded-xl bg-zinc-900 p-4">
            <div className="mb-2 flex items-center justify-between">
              <span className="font-mono text-sm">{f.control}</span>
              <span className={`rounded px-2 py-0.5 text-xs font-bold text-white ${sevColor[f.severity] ?? "bg-zinc-600"}`}>
                {f.severity}
              </span>
            </div>
            <p className="text-sm text-zinc-300">{f.evidence}</p>
            <p className="mt-2 text-xs text-zinc-500">
              Checked {new Date(f.checkedAt).toLocaleTimeString()} on {new Date(f.checkedAt).toLocaleDateString()}
            </p>
            {f.status !== "pass" && (
              <p className="mt-2 text-xs text-zinc-500">Fix: {f.remediation}</p>
            )}
          </div>
        ))}
      </div>
    </main>
  );
}