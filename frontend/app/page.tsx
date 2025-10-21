'use client';

import { useState, useEffect, FormEvent } from 'react';
import toast from 'react-hot-toast';
import { Loader2, Clipboard, Twitter, Linkedin, FileText } from 'lucide-react';

// 1. Define the Job type
interface Job {
  ID: number;
  original_url: string;
  status: string;
  status_detail: string;
  summary: string;
  tweets: string; 
  linkedin_post: string;
  CreatedAt: string;
}

// Helper component for job status 
const StatusBadge = ({ status }: { status: string }) => {
  let className = 'badge';
  if (status === 'complete') className = 'badge badge-success';
  if (status === 'failed') className = 'badge badge-error';
  if (status === 'processing') className = 'badge badge-info';
  if (status === 'pending') className = 'badge badge-ghost';

  return <div className={`${className} capitalize`}>{status}</div>;
};

// Main Page Component
export default function HomePage() {
  const [url, setUrl] = useState('');
  const [jobs, setJobs] = useState<Job[]>([]);
  const [isLoading, setIsLoading] = useState(false);

  const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

  // Function to fetch all jobs
  const fetchJobs = async () => {
    try {
      const response = await fetch(`${API_URL}/api/jobs`);
      if (!response.ok) throw new Error('Failed to fetch jobs');
      const data: Job[] = await response.json();
      setJobs(data);
    } catch (error) {
      console.error(error);
      toast.error('Could not fetch job history.');
    }
  };

  // Fetch jobs on load and set up polling (same as before)
  useEffect(() => {
    fetchJobs();
    const interval = setInterval(fetchJobs, 5000);
    return () => clearInterval(interval);
  }, []);

  // Handle form submission
  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setIsLoading(true);

    try {
      const response = await fetch(`${API_URL}/api/jobs`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ url }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Failed to submit job');
      }

      const newJob = await response.json();
      setJobs([newJob, ...jobs]);
      setUrl('');
      toast.success('Job Submitted! Generating content...');
    } catch (error: any) {
      toast.error(`Submission Failed: ${error.message}`);
    } finally {
      setIsLoading(false);
    }
  };

  // Helper to copy text to clipboard
  const copyToClipboard = (text: string, type: string) => {
    navigator.clipboard.writeText(text);
    toast.success(`${type} copied to clipboard!`);
  };

  // 11. Helper to parse and display tweets
  const renderTweets = (tweetsJson: string) => {
    try {
      const tweets: string[] = JSON.parse(tweetsJson);
      return (
        <ul className="space-y-3">
          {tweets.map((tweet, index) => (
            <li key={index} className="flex items-start gap-3">
              <Twitter className="h-4 w-4 mt-1 flex-shrink-0" />
              <span className="flex-grow">{tweet}</span>
              <button
                className="btn btn-ghost btn-square btn-sm"
                onClick={() => copyToClipboard(tweet, 'Tweet')}
              >
                <Clipboard className="h-4 w-4" />
              </button>
            </li>
          ))}
        </ul>
      );
    } catch (e) {
      return null;
    }
  };

  // The JSX for rendering
  return (
    <main className="container mx-auto max-w-4xl p-4 md:p-8">
      <header className="text-center mb-8">
        <h1 className="text-4xl md:text-5xl font-bold mb-2">Content-Genie 🤖</h1>
        <p className="text-lg text-base-content/70">
          Turn any article into a summary, tweets, and a LinkedIn post.
        </p>
      </header>

      {/* --- Submission Form --- */}
      <div className="card bg-base-200 shadow-xl mb-8">
        <div className="card-body">
          <form onSubmit={handleSubmit} className="flex flex-col sm:flex-row gap-2">
            <input
              type="url"
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              placeholder="Enter an article URL (e.g., https://...)"
              className="input input-bordered w-full"
              required
              disabled={isLoading}
            />
            <button
              type="submit"
              disabled={isLoading}
              className="btn btn-primary w-full sm:w-auto"
            >
              {isLoading ? (
                <>
                  <span className="loading loading-spinner loading-xs"></span>
                  Generating...
                </>
              ) : (
                'Generate Content'
              )}
            </button>
          </form>
        </div>
      </div>

      {/* --- Job List --- */}
      <section className="space-y-6">
        <h2 className="text-2xl font-semibold">Your Content History</h2>
        {jobs.length === 0 && !isLoading ? (
          <p className="text-center text-base-content/70">
            No jobs yet. Submit a URL to get started!
          </p>
        ) : (
          jobs.map((job) => (
            <div key={job.ID} className="card bg-base-200 shadow-md overflow-hidden">
              <div className="card-body">
                <div className="flex justify-between items-center mb-2">
                  <p className="font-mono text-sm text-base-content/60 truncate w-3/4">
                    {job.original_url}
                  </p>
                  <StatusBadge status={job.status} />
                </div>

                {job.status === 'processing' && (
                  <div className="flex items-center text-info text-sm">
                    <span className="loading loading-spinner loading-xs mr-2"></span>
                    {job.status_detail || 'Processing...'}
                  </div>
                )}

                {job.status === 'failed' && (
                  <p className="text-error text-sm">{job.status_detail}</p>
                )}

                {job.status === 'complete' && (
                  <div className="space-y-6 pt-4">
                    {/* Summary */}
                    <div className="space-y-2">
                      <h4 className="flex items-center font-semibold">
                        <FileText className="h-4 w-4 mr-2" /> Summary
                      </h4>
                      <p className="text-base-content/80 p-4 bg-base-300 rounded-md">
                        {job.summary}
                      </p>
                      <button
                        className="btn btn-outline btn-sm"
                        onClick={() => copyToClipboard(job.summary, 'Summary')}
                      >
                        <Clipboard className="h-4 w-4 mr-2" /> Copy Summary
                      </button>
                    </div>

                    {/* LinkedIn Post */}
                    <div className="space-y-2">
                      <h4 className="flex items-center font-semibold">
                        <Linkedin className="h-4 w-4 mr-2" /> LinkedIn Post
                      </h4>
                      <p className="text-base-content/80 p-4 bg-base-300 rounded-md whitespace-pre-wrap">
                        {job.linkedin_post}
                      </p>
                      <button
                        className="btn btn-outline btn-sm"
                        onClick={() =>
                          copyToClipboard(job.linkedin_post, 'LinkedIn Post')
                        }
                      >
                        <Clipboard className="h-4 w-4 mr-2" /> Copy Post
                      </button>
                    </div>

                    {/* Tweets */}
                    <div className="space-y-2">
                      <h4 className="flex items-center font-semibold">
                        <Twitter className="h-4 w-4 mr-2" /> Tweets
                      </h4>
                      <div className="text-base-content/80 p-4 bg-base-300 rounded-md">
                        {renderTweets(job.tweets)}
                      </div>
                    </div>
                  </div>
                )}
              </div>
            </div>
          ))
        )}
      </section>
    </main>
  );
}
