export interface Ingredient<T = number> extends IngredientDetails<T>{
  name: string,
}

export interface IngredientDetails<T=number> {
  amount: T,
  unit: string
}

export interface ExtendedRecipe extends Recipe {
  removed: boolean;
}

export interface Recipe {
  id: string
  title: string
  slug: string
  description: string
  ingredients: Ingredient[]
}

export interface ApiResponse<T = Object> {
  success: boolean
  data?: T
  message?: string
}

export enum LSKeys {
  RECIPES = 'recipes'
}