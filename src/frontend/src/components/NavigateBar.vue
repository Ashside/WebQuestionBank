<template>
  <nav class="navigate">
    <ul>
      <li v-for="item in items" :key="item.text" @click="navigate(item.link)">
        <a href="#">
          <span>{{ item.text }}</span>
          <i :class="item.icon"></i>
        </a>
      </li>
    </ul>
  </nav>
</template>

<script>
import router from "@/router";

export default {
  name: 'NavigateBar',
  props: {
    items: {
      type: Array,
      required: true,
      default: () => [
        { text: 'Home', link: '/', icon: 'fas fa-home' },
        { text: 'About', link: '/about', icon: 'fas fa-user' },
        { text: 'Services', link: '/services', icon: 'fas fa-cog' },
        { text: 'Contact', link: '/contact', icon: 'fas fa-envelope' },
      ],
    },
  },
  methods: {
    navigate(link) {
      this.$emit('navigate', link);
      router.push(link);
    },
  },
};
</script>

<style scoped>
.navigate {
  background: #1e1e1e;
  padding: 10px 20px;
}

.navigate ul {
  list-style: none;
  display: flex;
  justify-content: space-around;
  align-items: center;
  margin: 0;
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
}

.navigate a:hover {
  color: yellow;
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
