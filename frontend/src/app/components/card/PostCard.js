"use client";

import Link from 'next/link';
import styles from './PostCard.module.css'; 
import { forwardRef } from 'react';

const PostCard = forwardRef(({ postId, authorName, authorAvatar, content, image, likes, comments, createdAt, type }, ref) => {
  const formattedDate = new Date(createdAt).toLocaleString('en-US', {
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  });
  const getTypeDisplay = (type) => {
    if (type === 'pb') return 'Public';
    if (type === 'pv') return 'Private';
    if (type === 'ap') return 'Almost Private';
    return 'Unknown'; // Fallback
  };

  const  handleLike = async (postID) =>{
    const path = "/api/post/like";
    const body = {
      post_id: parseInt(postID),
      user_id: Number(localStorage.getItem('id'))
    };

    // const data = await fetcher.post(path, body);
    // if (data && data.msg) {
    //   return;
    // }
    // await getPost(postID);
  }

  return (
    <div ref={ref} className={styles.post_card}>
      <h1>
      {type}
      </h1>
      <div  className={styles.author}>
        <img src={authorAvatar} alt={`${authorName}'s avatar`}/>
        <div className={styles.author_info}>
          <p className={styles.author_name}>
            <Link href={`/profile/${authorName.toLowerCase().replace(' ', '-')}`}>
              {authorName}
            </Link>
          </p>
          <p className={styles.post_info}>
            <span className={styles.type}>{getTypeDisplay(type)} -</span> {formattedDate}
          </p>
        </div>
      </div>
      <div className={styles.post}>
        <p className={styles.post_content}>{content}</p>
         <img src={image} alt="image" className={styles.post_image} />
      </div>
      <div className={styles.reactions}>
        <button className={styles.reaction} type="button" name="reaction" value="like" onClick={() => handleLike(postID)}>
          <i className="fa-solid fa-thumbs-up"></i> {likes.toLocaleString()}
        </button>
        <button className={styles.reaction} type="button" name="comments" value="comments">
          <i className="fa-solid fa-message"></i> {comments?.toLocaleString() || 0}
        </button>
      </div>
    </div>
  );
});

PostCard.displayName = 'PostCard';

export default PostCard;