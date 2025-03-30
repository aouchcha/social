"use client";

import { useState } from 'react';
import styles from './CreatePost.module.css';

const CreatePost = () => {
  const [content, setContent] = useState('');
  const [type, setType] = useState("pb");
  const [file, setFile] = useState(null);
  const [error, setError] = useState('');

  const MAX_CONTENT_LENGTH = 1000;
  const MAX_FILE_SIZE = 20 * 1024 * 1024;
  const ALLOWED_FILE_TYPES = ['image/png', 'image/jpeg', 'image/gif'];

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');

    // Validation
    if (!content.trim() || content.length > MAX_CONTENT_LENGTH) {
      setError(`Content must not be between 1 to ${MAX_CONTENT_LENGTH} characters.`);
      return;
    }
    
    if (!["pb", "pv", "ap"].includes(type)) {
      setError('Invalid post type.');
      return;
    }

    let imageBase64 = '';
    if (file) {
      if (!ALLOWED_FILE_TYPES.includes(file.type)) {
        setError('Only .png, .jpeg, and .gif files are allowed.');
        return;
      }
      if (file.size > MAX_FILE_SIZE) {
        setError('File size must not exceed 20MB.');
        return;
      }

      const reader = new FileReader();
      reader.readAsDataURL(file);
      imageBase64 = await new Promise((resolve) => {
        reader.onload = () => resolve(reader.result);
      });
    }

    // data
    const formData ={
        user_id: Number(JSON.parse(localStorage.getItem('user')).userId),
        content: content,
        privacy: type,
        image_url: imageBase64,
    }

    try {
      const response = await fetch('http://localhost:8080/api/post/create', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        throw new Error('Failed to create post');
      }

      const result = await response.json();
      console.log('Post created:', result);

      setContent('');
      setType("pb");
      setFile(null);
    } catch (err) {
      setError(err.message || 'Something went wrong.');
    }
  };

  return (
    <div>
      <form className={styles.create_post} onSubmit={handleSubmit}>
        <div className={styles.textarea}>
          <img
            src="https://images.unsplash.com/photo-1614204424926-196a80bf0be8?q=80&w=1374&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
            alt="User avatar"
            className={styles.avatar}
          />
          <textarea
            id="textarea"
            name="content"
            rows="4"
            cols="50"
            placeholder="Write your post here..."
            value={content}
            onChange={(e) => setContent(e.target.value)}
            maxLength={MAX_CONTENT_LENGTH}
            required
          />
        </div>
        {error && <p className={styles.error}>{error}</p>}
        <div className={styles.btns}>
          <div className={styles.options}>
            <select
              name="type"
              value={type}
              onChange={(e) => setType(e.target.value)}
              className={styles.select}
            >
              <option value="pb">Public</option>
              <option value="pv">Private</option>
              <option value="ap">Almost Private</option>
            </select>
            <input
              type="file"
              name="image"
              id="upload"
              accept=".png,.jpeg,.jpg,.gif"
              onChange={(e) => setFile(e.target.files[0])}
              className={styles.fileInput}
            />
          </div>
          <button type="submit" className={styles.post}>
            Post
          </button>
        </div>
      </form>
    </div>
  );
};

export default CreatePost;