<template>
  <div>
    <Home v-if="currentPage === 'home'" @selectVideo="viewVideo" />
    <VideoDetail
      v-if="currentPage === 'videoDetail'"
      :videoId="selectedVideoId"
      @goHome="goHome"
      @rateVideo="rateVideo"
    />
    <Rating
      v-if="currentPage === 'rating'"
      :videoId="selectedVideoId"
      @goHome="goHome"
    />
  </div>
</template>

<script>
import Home from './components/Home.vue';
import VideoDetail from './components/VideoDetail.vue';
import Rating from './components/Rating.vue';

export default {
  components: {
    Home,
    VideoDetail,
    Rating
  },
  data() {
    return {
      currentPage: 'home',
      selectedVideoId: null
    };
  },
  methods: {
    viewVideo(videoId) {
      this.selectedVideoId = videoId;
      this.currentPage = 'videoDetail';
    },
    goHome() {
      this.currentPage = 'home';
    },
    rateVideo(videoId) {
      this.selectedVideoId = videoId;
      this.currentPage = 'rating';
    }
  }
};
</script>