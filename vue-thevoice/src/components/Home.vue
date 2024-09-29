<template>
    <div>
      <h1>视频列表</h1>
      <ul>
        <li v-for="video in videoConfig" :key="video.videoId">
          <h3>{{ video.title }}</h3>
          <button @click="$emit('selectVideo', video.videoId)">查看详情</button>
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  import { videoConfig } from '../config/videoConfig';
  
  export default {
    data() {
      return {
        videoConfig
      };
    }
  };
  </script>