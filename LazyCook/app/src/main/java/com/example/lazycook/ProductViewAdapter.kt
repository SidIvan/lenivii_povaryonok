package com.example.lazycook

import android.app.Activity
import android.content.Context
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ArrayAdapter
import android.widget.ImageView
import android.widget.TextView

class ProductViewAdapter(context: Context, items: ArrayList<String>) :
    ArrayAdapter<String>(context, R.layout.productlist_row, items) {
    private var list: ArrayList<String> = items

    override fun getView(position: Int, convertView: View?, parent: ViewGroup): View {
        if (convertView == null) {
            var layoutInflater: LayoutInflater = LayoutInflater.from(context)
            var convertView2: View = layoutInflater.inflate(R.layout.productlist_row, null)

            var number: TextView = convertView2.findViewById(R.id.number)
            number.text = (position + 1).toString()

            var name: TextView = convertView2.findViewById(R.id.name)
            name.text = list[position]

            var remove : ImageView = convertView2.findViewById(R.id.delete_button)
            remove.setOnClickListener(object : View.OnClickListener {
                override fun onClick(v: View?) {
                    MainActivity.removeProductItem(position)
                }
            })
            return convertView2
        } else {
            return convertView
        }
    }
}