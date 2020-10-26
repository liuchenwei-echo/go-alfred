package alfred

func BuildTextItem(title string, subTitle string, arg string) Item {
	return Item{
		Uid: title,
		Subtitle: subTitle,
		Arg: arg,
		Type: "text",
	}
}

func BuildItemList(items []Item) ItemList {
	return ItemList{
		Items: items,
	}
}

func (l ItemList) Append(item Item) {
	l.Items = append(l.Items, item)
}

func EmptyItemList() ItemList {
	return ItemList{}
}
