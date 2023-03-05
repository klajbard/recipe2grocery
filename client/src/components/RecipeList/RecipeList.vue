<script setup lang="ts">
import axios from "axios";
import { onMounted, reactive, ref } from "vue";
import { ApiResponse, Recipe } from "../../models";
import RecipeItem from "./RecipeItem.vue";
let pageNum = 1;

const recipes = reactive<Recipe[]>([]);
const reachedEnd = ref(false);

const loadMore = async () => {
  axios
    .get<ApiResponse<Recipe[]>>(
      `${import.meta.env.VITE_BACKEND_URL}/recipes/${pageNum}`
    )
    .then((resp) => {
      if (resp.data == null || !resp.data.data) {
        throw new Error("No more data to fetch");
      }
      recipes.push(...resp.data.data);
      pageNum += 1;
    })
    .catch((err) => {
      console.error(err);
      if (err.message === "No more data to fetch") {
        reachedEnd.value = true;
      }
    });
};

onMounted(() => {
  loadMore();
});
</script>
<template>
  <h1>Recipes</h1>
  <div class="recipes">
    <recipe-item v-for="recipe in recipes" :key="recipe.slug" :recipe="recipe">
    </recipe-item>
    <button v-if="!reachedEnd" @click="loadMore">Load more</button>
    <div class="empty" v-if="reachedEnd">No more recipe to load...</div>
  </div>
</template>
<style scoped>
.recipes {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.empty {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0.6rem 1rem;
  border-radius: 8px;
  background: var(--button-background);
}
</style>
