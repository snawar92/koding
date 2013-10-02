class ChatConversationListView extends KDListView

  constructor:(options = {}, data)->

    options.itemClass  = ChatConversationListItem
    options.cssClass   = KD.utils.curry "chat-list", options.cssClass
    options.tagName    = "ul"

    super options, data

  goUp:(item)->
    index = @getItemIndex item
    return unless index >= 0

    if index - 1 >= 0
      item.collapseConversation()
      @items[index - 1].expandConversation() # toggleConversation()

  goDown:(item)->
    index = @getItemIndex item
    return unless index >= 0

    if index + 1 < @items.length
      item.collapseConversation()
      @items[index + 1].expandConversation() # toggleConversation()
