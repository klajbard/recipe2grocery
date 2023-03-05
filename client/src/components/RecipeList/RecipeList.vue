<script setup lang="ts">
import axios from "axios";
import { onMounted, reactive, ref } from "vue";
import { ApiResponse, Recipe } from "../../models";
import { recipesStore } from "../../stores/recipes.store";
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
    <div class="recipe" v-for="recipe in recipes" :key="recipe.slug">
      <div class="content">
        <h3 class="title">{{ recipe.title }}</h3>
        <div class="box">
          <ul class="ingredients">
            <li v-for="ingredient in recipe.ingredients" :key="ingredient.name">
              {{ ingredient.name }}
            </li>
          </ul>
        </div>
      </div>
      <button
        class="add-remove-button"
        @click="
          recipesStore.recipes.find(({ slug }) => slug === recipe.slug)
            ? recipesStore.removeRecipe(recipe)
            : recipesStore.addRecipe(recipe)
        "
      >
        {{
          recipesStore.recipes.find(({ slug }) => slug === recipe.slug)
            ? "-"
            : "+"
        }}
      </button>
    </div>
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

.recipe {
  padding: 0.5rem 1rem;
  background: #f39f8644;
  border-radius: var(--border-radius);
  text-align: left;
  gap: 1rem;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}

.title {
  margin-top: 0.5rem;
  margin-bottom: 0.5rem;
}

.description {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-align: justify;
}

.add-remove-button {
  width: 4rem;
}

.ingredients {
  box-sizing: border-box;
  display: flex;
  flex-wrap: wrap;
  padding: 0;
  margin: 0;
  justify-content: space-evenly;
  list-style: none;
  gap: 0.5rem;
}

.empty {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0.6rem 1rem;
  border-radius: 8px;
  background: var(--button-background);
}

.ingredients li {
  background: #fff1;
  border-radius: var(--border-radius);
  padding: 0.5rem 1rem;
  cursor: default;
}
</style>
