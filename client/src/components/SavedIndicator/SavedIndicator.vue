<script setup lang="ts">
import { recipesStore } from "../../stores/recipes.store";
</script>
<template>
  <router-link :to="{ name: 'Saved Recipes' }" class="recipes">
    {{ recipesStore.recipes.length }} ✔
  </router-link>
</template>
<style scoped>
.recipes {
  position: fixed;
  top: 2rem;
  right: 2rem;
  height: 3rem;
  padding: 0 0.5rem;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #94c99744;
  border-radius: var(--border-radius);
  color: white;
}
</style>
