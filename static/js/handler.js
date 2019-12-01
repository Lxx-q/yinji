
window.AUDIO_HANDLER = {
	receive:function( item ){
		var newItem = {};

		newItem = item;

		var image = newItem.image;

		if( image == "" ){
			image = "resources/image/none.jpg";
		}



		newItem.image = "/yinji" + "/" + image;

		return newItem;
	}
}