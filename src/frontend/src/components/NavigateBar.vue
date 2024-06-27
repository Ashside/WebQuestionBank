<template>
  <nav class="navigate">
    <ul>
      <li v-for="item in items" :key="item.text" @click="navigate(item.link, item.text)">
        <a href="#">
          <span>{{ item.text }}</span>
          <i :class="item.icon"></i>
        </a>
      </li>
    </ul>
  </nav>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  name: 'NavigateBar',
  props: {
    items: {
      type: Array,
      required: true,
      default: () => [
        { text: 'Home', link: '/home', icon: 'fas fa-home' },
        { text: 'About', link: '/about', icon: 'fas fa-user' },
        { text: 'Logout', link: '/', icon: 'fas fa-sign-out-alt' },
        { text: 'Contact', link: '/contact', icon: 'fas fa-envelope' },
      ],
    },
  },
  methods: {
    ...mapActions(['logout']),
    navigate(link, text) {
      if (text === 'Logout') {
        this.logout();
      }
      this.$emit('navigate', link);
      this.$router.push(link);
    },
  },
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Dancing+Script:wght@400;700&display=swap');

.navigate {
  background-image: linear-gradient(to right, #333333, #5a6268);
  padding: 1px 20px;
}

.navigate ul {
  list-style: none;
  display: flex;
  justify-content: space-around;
  align-items: center;
  margin: 0 auto; /* 居中显示 */
  padding: 0;
}

.navigate li {
  position: relative;
  overflow: hidden;
}

.navigate a {
  display: flex;
  align-items: center;
  color: #fff;
  text-decoration: none;
  padding: 10px;
  transition: color 0.3s ease;
  font-family: 'Dancing Script', cursive; /* 设置花体字体 */
  font-size: 1.5em;
}

.navigate a:hover {
  color: #d4edda;
}

.navigate span {
  margin-right: 10px;
}

.navigate i {
  transition: transform 0.3s ease;
}

.navigate a:hover i {
  transform: scale(1.2);
}

.navigate li::before {
  content: "";
  position: absolute;
  background: #fff;
  height: 100%;
  width: 100%;
  top: 100%;
  left: 0;
  transition: all 0.3s ease;
  z-index: -1;
}

.navigate li:hover::before {
  top: 0;
}
</style>
