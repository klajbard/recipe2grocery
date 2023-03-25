<script setup lang="ts">
import axios from "axios";
import { ref, watch } from "vue";
import { ApiResponse, ExtendedRecipe } from "../../models";
import { recipesStore } from "../../stores/recipes.store";
import IconElement from "../Icons/IconElement.vue";

const props = defineProps<{ recipe: ExtendedRecipe }>();
defineEmits(['show-overlay'])

const isAdded = ref<boolean>(!!recipesStore.recipes.find(
  ({ slug }) => slug === props.recipe.slug
));
const container = ref<HTMLDivElement | null>(null);
watch(recipesStore.recipes, () => {
  isAdded.value = !!recipesStore.recipes.find(
    ({ slug }) => slug === props.recipe.slug
  );
});

const handleUndo = () => {
  axios
    .post<ApiResponse>(`${import.meta.env.VITE_BACKEND_URL}/recipe?id=${props.recipe.id}`, {
      title: props.recipe.title,
      description: props.recipe.description,
      ingredients: props.recipe.ingredients,
    })
    .then((resp) => {
      if (resp.data.success) {
        props.recipe.removed = false;
      }
    });
};

</script>
<template>
  <div v-if="recipe.removed">
    <button class="undo" @click="handleUndo()">
      Undo removal of {{ recipe.title }}
    </button>
  </div>
  <div v-else class="recipe" ref="container" :class="isAdded && 'selected'">
    <h3 class="title">{{ recipe.title }}</h3>
    <div class="button-group">
      <button class="info" @click="$emit('show-overlay')" :title="`More information about ${recipe.title}`"><icon-element
          name="info" /></button>
      <button class="add" @click="
        isAdded
          ? recipesStore.removeRecipe(recipe)
          : recipesStore.addRecipe(recipe)
      " :title="isAdded ? `Remove ${recipe.title} from list` : `Add ${recipe.title} to list`">
        <icon-element :name="isAdded ? 'minus' : 'plus'" />
      </button>
    </div>
  </div>
</template>
<style scoped>
.recipe {
  padding: 0.75rem 1rem;
  background: var(--background-brown);
  border-radius: var(--border-radius);
  text-align: left;
  gap: 0.5rem;
  display: flex;
  justify-content: space-between;
}

.recipe.selected {
  background: var(--dark-green);
}

.title {
  margin: 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 100%;
}

.button-group {
  display: flex;
  gap: 0.5rem;
  align-items: center;
}

.button-group button {
  height: 2.5rem;
  min-width: 4rem;
  flex: 1;
  gap: 0.5rem;
  font-size: 0.75rem;
  white-space: nowrap;
}

.add {
  background-color: var(--color-green-fade);
}

.info {
  background-color: var(--color-blue-fade);
}

.undo {
  width: 100%;
}
</style>
