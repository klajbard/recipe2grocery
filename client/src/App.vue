<script setup lang="ts">
import { onMounted } from "vue";
import Indicator from "./components/SavedIndicator/SavedIndicator.vue";
import {recipesStore} from './stores/recipes.store'

onMounted(() => {
  recipesStore.readLS()
});
</script>

<template>
  <header>
    <router-link :to="{ name: 'Recipes' }">Recipes</router-link> |
    <router-link :to="{ name: 'Add Recipe' }">Add Recipe</router-link> |
    <router-link :to="{ name: 'Saved Recipes' }">Saved Recipes</router-link>
  </header>

  <indicator />
  <router-view />
</template>

<style scoped>
header {
  padding: 1rem;
}
</style>
