package com.example.lazycook

import android.Manifest
import android.content.Intent
import android.content.pm.PackageManager
import android.graphics.Bitmap
import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.provider.MediaStore
import android.view.View
import android.widget.AdapterView
import android.widget.ArrayAdapter
import android.widget.EditText
import android.widget.ImageView
import android.widget.ListView
import android.widget.Toast
import androidx.core.content.ContextCompat

class MainActivity : AppCompatActivity() {
    companion object {
        private var items: ArrayList<String> = ArrayList()
        private lateinit var adapter: ArrayAdapter<String>

        fun addProductItem(item: String) {
            items.add(item)
            adapter.notifyDataSetChanged()
        }

        fun removeProductItem(index: Int) {
            items.removeAt(index)
            adapter.notifyDataSetChanged()
        }
    }

    private val PERMISSION_CODE = 1000
    private val IMAGE_CAPTURE_CODE = 1001
    private lateinit var productView: ListView
    private lateinit var inputLine: EditText
    private lateinit var enterText: ImageView
    private lateinit var cameraButton: ImageView

    override fun onActivityResult(requestCode: Int, resultCode: Int, data: Intent?) {
        super.onActivityResult(requestCode, resultCode, data)
        if (requestCode == IMAGE_CAPTURE_CODE) {
            val image: Bitmap = data?.extras?.get("data") as Bitmap
        }
    }

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        productView = findViewById(R.id.listview)
        inputLine = findViewById(R.id.input)
        enterText = findViewById(R.id.add)
        adapter = ArrayAdapter(applicationContext, android.R.layout.simple_list_item_1, items)
        productView.adapter = adapter
        cameraButton = findViewById(R.id.camera)

        enterText.setOnClickListener(object : View.OnClickListener {
            override fun onClick(view: View?) {
                val text: String = inputLine.text.toString()
                if (text == null || text.isEmpty()) {
                    Toast.makeText(
                        applicationContext,
                        getString(R.string.ErrorEmptyInput),
                        Toast.LENGTH_SHORT
                    ).show()
                } else {
                    addProductItem(text)
                    inputLine.setText("")
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

        cameraButton.setOnClickListener(object : View.OnClickListener {
            override fun onClick(v: View?) {
                var permissions: Array<String> = arrayOf()
                if (ContextCompat.checkSelfPermission(
                        applicationContext,
                        Manifest.permission.CAMERA
                    ) == PackageManager.PERMISSION_DENIED) {
                    permissions += Manifest.permission.CAMERA
                }
                if (permissions.isNotEmpty()) {
                    requestPermissions(permissions, PERMISSION_CODE)

                }

                val cameraIntent = Intent(MediaStore.ACTION_IMAGE_CAPTURE)
                startActivityForResult(cameraIntent, IMAGE_CAPTURE_CODE)
            }
        })
    }
}