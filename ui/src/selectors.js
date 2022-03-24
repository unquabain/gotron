import { createSelector } from '@reduxjs/toolkit'

export const sessionSelector = state => (state || {}).session || {}
export const dataSelector = state => (state || {}).data || {}
export const viewsSelector = state => (state || {}).views || {}

export const usernameSelector = createSelector(sessionSelector, session => session.username)
export const navPathSelector = createSelector(sessionSelector, session => session.navPath)

export const fortuneSelector = createSelector(viewsSelector, views => views.fortune || '')
export const catfactSelector = createSelector(viewsSelector, views => views.catfact || '')
