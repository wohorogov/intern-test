import React from 'react';
import SearchTips from '../components/search-tips';

const SearchForm =
  () =>
    (
      <section>
        <h1>Поисковые подсказки</h1>
        <input className="search-input" type="search"/>
        <SearchTips />
      </section>
    );

export default SearchForm;
