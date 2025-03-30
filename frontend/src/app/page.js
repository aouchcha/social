"use client";

import { useCallback, useEffect, useRef, useState } from 'react';
import PostCard from './components/card/PostCard';
import Style from './page.module.css'
import CreatePost from './components/PostCreation/CreatePost';
import { Header } from './components/headerComp';

export default function Home() {
    const [posts, setPosts] = useState([]);
    const [offset, setOffset] = useState(0);
    const [hasMore, setHasMore] = useState(true);
    const [loading, setLoading] = useState(false);
    const observer = useRef(null);
    const limit = 10;

    const lastPostRef = useCallback((node) => {
        if (loading) return;
        if (observer.current) observer.current.disconnect();
        observer.current = new IntersectionObserver((entries) => {
          if (entries[0].isIntersecting && hasMore) {
            setOffset((prevOffset) => prevOffset + limit);
          }
        });
        if (node) observer.current.observe(node);
      }, [loading, hasMore]);
    
      const fetchPosts = useCallback(async () => {
        if (!hasMore || loading) return;
        setLoading(true);
    
        try {
          const response = await fetch(`http://localhost:8080/api/posts/?limit=${limit}&offset=${offset}`);
          if (!response.ok) throw new Error('Failed to fetch posts');
          const newPosts = await response.json();
    
          if (newPosts.length < limit) {
            setHasMore(false);
          }
    
          setPosts((prevPosts) => {
            const existingIds = new Set(prevPosts.map((post) => post.postId));
            const uniqueNewPosts = newPosts.filter((post) => !existingIds.has(post.postId));
            return [...prevPosts, ...uniqueNewPosts];
          });
        } catch (err) {
          console.error(err);
        } finally {
          setLoading(false);
        }
      }, [offset, hasMore, loading]);
    
      useEffect(() => {
        fetchPosts();
      }, [offset, fetchPosts]);
    
      const handlePostCreated = (newPost) => {
        setPosts((prevPosts) => {
          if (prevPosts.some((post) => post.postId === newPost.postId)) {
            return prevPosts;
          }
          return [newPost, ...prevPosts]; 
        });
      };

    return (
      <>
      <Header/>
        <div className={Style.container}>
            <div className={Style.item}>Header</div>
            <div className={Style.item}>Sidebar</div>
            <div className={Style.item}>
                <CreatePost onPostCreated={handlePostCreated} />
                {posts.map((post, index) => {
                    return(
                    <PostCard
                        ref={index === posts.length - 1 ? lastPostRef : null}
                        key={post.post_id}
                        postId={post.post_id}
                        authorName={post.username}
                        authorAvatar={`https://robohash.org/${post.username}`}
                        content={post.content}
                        image={`http://localhost:8080/${post.image_url}`}
                        likes={post.likes}
                        comments={post.comment_count}
                        createdAt={post.created_at}
                        type={post.privacy}
                    />
                );
                })}
            {loading && <p>Loading more posts...</p>}
            </div>
            <div className={Style.item}>Sidebar</div>
        </div>
        </>
    );
}