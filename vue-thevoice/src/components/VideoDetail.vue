<template>
    <div>
      <h2>{{ video.title }}</h2>
      <div>
        <button @click="playOriginal">原片播放</button>
        <button @click="playDubbed">配音播放</button>
        <button @click="$emit('rateVideo', video.videoId)">评分</button>
        <button @click="$emit('goHome')">返回主页</button>
      </div>
      <video v-if="playing === 'original'" controls :src="video.originalVideoUrl">原视频加载失败</video>
      <video v-if="playing === 'dubbed'" controls :src="video.dubbedVideoUrl">配音视频加载失败</video>
    </div>
  </template>
  
  <script>
  import { videoConfig } from '../config/videoConfig';
  
  export default {
    props: ['videoId'],
    data() {
      return {
        video: null,
        playing: null
      };
    },
    created() {
      this.video = videoConfig.find(v => v.videoId === this.videoId);
    },
    methods: {
      playOriginal() {
        this.playing = 'original';
      },
      playDubbed() {
        this.playing = 'dubbed';
      }
    }
  };
  </script>