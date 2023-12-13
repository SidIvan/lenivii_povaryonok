package com.example.lazycook

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.view.View
import android.widget.AdapterView
import android.widget.ArrayAdapter
import android.widget.EditText
import android.widget.ImageView
import android.widget.ListView
import android.widget.Toast

class MainActivity : AppCompatActivity() {
    private lateinit var productView: ListView
    private var items: ArrayList<String> = ArrayList()
    private lateinit var adapter: ArrayAdapter<String>
    private lateinit var input: EditText
    private lateinit var enter: ImageView

    private fun addProductItem(item: String) {
        items.add(item)
        adapter.notifyDataSetChanged()
    }

    private fun removeProductItem(index: Int) {
        items.removeAt(index)
        adapter.notifyDataSetChanged()
    }

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        productView = findViewById(R.id.listview)
        input = findViewById(R.id.input)
        enter = findViewById(R.id.add)
        adapter = ArrayAdapter(applicationContext, android.R.layout.simple_list_item_1, items)
        productView.adapter = adapter

        enter.setOnClickListener(object : View.OnClickListener {
            override fun onClick(view: View?) {
                val text: String = input.text.toString()
                if (text == null || text.isEmpty()) {
                    val toast = Toast.makeText(
                        applicationContext,
                        getString(R.string.ErrorEmptyInput),
                        Toast.LENGTH_SHORT
                    )
                    toast.show()
                } else {
                    addProductItem(text)
                    input.setText("")
                }
            }
        })

        productView.setOnItemLongClickListener(object : AdapterView.OnItemLongClickListener {
            override fun onItemLongClick(
                parent: AdapterView<*>?,
                view: View?,
                position: Int,
                id: Long
            ): Boolean {
                removeProductItem(position)
                return false
            }
        })
    }
}