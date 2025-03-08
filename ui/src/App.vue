<template>
  <header class="main-app-bar">
    <MainAppBar @toggle="toggleMenu" />
  </header>

  <nav>
    <RouterLink to="/">Home</RouterLink>
    <RouterLink to="/manage">Manage</RouterLink>
    <RouterLink to="/me">Profile</RouterLink>
  </nav>

  <main>
    <RouterView />
  </main>
</template>

<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import MainAppBar from './components/MainAppBar.vue'

const toggleMenu = (isOpen: boolean) => {
  const nav = document.querySelector('nav')
  nav?.classList.toggle('close', !isOpen)
}
</script>

<style scoped>
header {
  background-color: var(--color-background);
  box-shadow: 0 -5px 10px rgba(0, 0, 0, 0.1);
}

nav {
  width: 100%;
  height: 3rem;
  font-size: 12px;
  text-align: center;
  overflow: hidden;
  transition: all 0.3s ease;

  background-color: var(--color-background);
  box-shadow: 0 -5px 10px rgba(0, 0, 0, 0.1);
}

nav.close {
  height: 0;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0.8rem 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

main {
  padding-bottom: 1.5rem;
}
</style>
