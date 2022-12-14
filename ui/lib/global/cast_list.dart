class CastList<From, To> extends Iterable<To> {
  final Iterable<From> _items;
  final To Function(From) _cast;

  CastList(this._items, this._cast);

  @override
  Iterator<To> get iterator =>
      CastListIterator<From, To>(_items.iterator, _cast);
}

class CastListIterator<From, To> implements Iterator<To> {
  final Iterator<From> _iterator;
  final To Function(From) _cast;

  CastListIterator(this._iterator, this._cast);

  @override
  To get current {
    return _cast(_iterator.current);
  }

  @override
  bool moveNext() {
    return _iterator.moveNext();
  }
}
