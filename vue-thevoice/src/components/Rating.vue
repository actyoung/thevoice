<template>
    <div>
      <h2>{{ video.title }}</h2>
  
      <h3>为角色评分</h3>
      <div v-for="(character, index) in video.characters" :key="index">
        <p>{{ character.name }}</p>
        <label>评分:</label>
        <select v-model="character.rating">
          <option v-for="n in 5" :key="n" :value="n">{{ n }}</option>
        </select>
      </div>
      <button @click="submitRatings">提交评分</button>
      <button @click="$emit('goHome')">返回主页</button>
  
      <div v-if="showRankings">
        <h3>角色排名</h3>
        <ul>
          <li v-for="(character, index) in sortedCharacters" :key="index">
            {{ character.name }}: {{ character.rating }}
          </li>
        </ul>
      </div>
    </div>
  </template>
  
  <script>
  import { videoConfig } from '../config/videoConfig';
  
  export default {
    props: ['videoId'],
    data() {
      return {
        video: null,
        showRankings: false
      };
    },
    computed: {
      sortedCharacters() {
        return this.video.characters.slice().sort((a, b) => b.rating - a.rating);
      }
    },
    methods: {
      submitRatings() {
        if (this.video.characters.every(character => character.rating !== null)) {
          this.showRankings = true;
          // 在这里可以添加提交评分到服务器的逻辑
        } else {
          alert('请为所有角色评分');
        }
      }
    },
    created() {
      this.video = videoConfig.find(v => v.videoId === this.videoId);
    }
  };
  </script>