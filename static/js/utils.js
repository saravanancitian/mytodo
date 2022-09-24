function formtojson(frm){
    const inputs = frm.elements;
    const data = {};
    for(let i = 0 ; i < inputs.length; i++){
        if(inputs[i].type == 'text'){
            let key = inputs[i].id;
            let val = inputs[i].value;
            data[key] = val;
        }
    }

    return data;
}