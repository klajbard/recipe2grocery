<script setup lang="ts">
import { recipesStore } from "../../stores/recipes.store";
import IconElement from "../Icons/IconElement.vue";
</script>
<template>
  <router-link :to="{ name: 'Saved Recipes' }" class="recipes">
    <icon-element name="cart" />
    {{ recipesStore.recipes.length }}
  </router-link>
</template>
<style scoped>
.recipes {
  position: absolute;
  top: 1rem;
  right: 0.5rem;
  height: 3rem;
  padding: 0 0.5rem;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #94c99744;
  border-radius: var(--border-radius);
  color: white;
  gap: 0.5rem;
  font-size: 1.25rem;
}
</style>
