<script setup lang="ts">
import { recipeStore } from "../../stores/recipe.store";
const units = [
  "gramm",
  "dkg",
  "kg",
  "liter",
  "dl",
  "ml",
  "cup",
  "tablespoon",
  "teaspoon",
  "bunch",
  "pinch",
  "pcs",
];
</script>
<template>
  <div class="form-section" style="background: #9fa4c444">
    <span>Ingredients</span>
    <div
      v-for="(ingredient, index) in recipeStore.ingredients"
      :key="index"
      class="form-group"
    >
      <div class="ingredient">
        <label :for="`ingredient-name-${index}`">Name</label>
        <label :for="`ingredient-amount-${index}`">Amount</label>
        <label :for="`ingredient-unit-${index}`">Unit</label>
        <input
          :id="`ingredient-name-${index}`"
          v-model="ingredient.name"
          name="ingredient-name"
        />
        <input
          :id="`ingredient-amount-${index}`"
          v-model.number="ingredient.amount"
          name="ingredient-amount"
        />
        <select
          :id="`ingredient-unit-${index}`"
          v-model="ingredient.unit"
          name="ingredient-unit"
        >
          <option v-for="unit in units" :value="unit" :key="unit">
            {{ unit }}
          </option>
        </select>
      </div>
      <button
        v-if="recipeStore.ingredients.length > 1"
        type="button"
        class="close"
        @click="recipeStore.removeIngredient(index)"
      >
        X
      </button>
    </div>
    <button
      v-if="
        recipeStore.ingredients[recipeStore.ingredients.length - 1].name !== ''
      "
      class="full-width"
      type="button"
      @click="recipeStore.addIngredient()"
    >
      Add ingredient
    </button>
  </div>
</template>
<style scoped>
.ingredient {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
}

.close {
  padding: 0.5rem 0.75rem;
  background: #ff000044;
}

.form-group {
  display: flex;
  gap: 0.5rem;
  align-items: start;
  padding: 0.5rem;
  background: #1c8c7344;
  border-radius: 4px;
}
</style>
