
window.AUDIO_HANDLER = {
	receive:function( item ){
		var newItem = {};

		newItem = item;

		var image = newItem.image;

		if( image == "" ){
			image = "resources/image/none.jpg";
		}

		var modifyTime = this.parseTime( newItem.modifyTime );
		var createTime = this.parseTime( newItem.createTime );

		newItem.modifyTime = modifyTime;
		newItem.createTime = createTime;

		newItem.image = "/yinji" + "/" + image;

		return newItem;
	},parseTime:function( date ){
		var arr=date.split("T");
     	var d=arr[0];
     	var darr = d.split('-');
     	var t=arr[1];
     	var tarr = t.split('.000');
     	var marr = tarr[0].split(':');
     	var dd = parseInt(darr[0])+"/"+parseInt(darr[1])+"/"+parseInt(darr[2])+" "+parseInt(marr[0])+":"+parseInt(marr[1])+":"+parseInt(marr[2]);
     	return dd;
	}
}