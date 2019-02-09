package io.gitee.stevenzack.foxui.FObject.Widget;

import android.support.v4.view.ViewCompat;
import android.view.View;
import android.widget.AdapterView;
import android.widget.ArrayAdapter;
import android.widget.Spinner;

import org.json.JSONArray;
import org.json.JSONTokener;

import java.util.ArrayList;
import java.util.List;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

public class FSpinner extends FObject {
    public Spinner v;
    private List<String> data_list = new ArrayList<>();
    private ArrayAdapter<String> adapter =null;
    private boolean selectedByCode;
    public FSpinner(FoxActivity foxActivity) {
        parentController = foxActivity;
        v = new Spinner(parentController);
        v.setId(ViewCompat.generateViewId());
        view=v;
        adapter=new ArrayAdapter<>(parentController,android.R.layout.simple_spinner_dropdown_item,data_list);
        v.setAdapter(adapter);
    }

    @Override
    public String getAttr(String attr) {
        String str = super.getAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr) {
            case "SelectedIndex":
                return String.valueOf(v.getSelectedItemPosition());
        }
        return "";
    }

    @Override
    public String setAttr(String attr, final String value, String value2) {
        String str = super.setAttr(attr, value, value2);
        if (str != null) {
            return str;
        }
        switch (attr) {
// -------------------------------------------------------------------
            case "NotifyDataSetChanged":
                adapter.notifyDataSetChanged();
                break;
            case "List":
                try {
                    JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    data_list.clear();
                    for (int i=0;i<array.length();i++) {
                        data_list.add(array.getString(i));
                    }
                    selectedByCode = true;
                    adapter.notifyDataSetChanged();
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "ListAdd":
                data_list.add(value);
                selectedByCode = true;
                adapter.notifyDataSetChanged();
                break;
            case "ListRemove":
                data_list.remove(Integer.parseInt(value));
                selectedByCode = true;
                adapter.notifyDataSetChanged();
                break;
            case "OnItemClick":
                v.setOnItemSelectedListener(new AdapterView.OnItemSelectedListener() {
                    @Override
                    public void onItemSelected(AdapterView<?> parent, View view, int position, long id) {
//                        Faithdroid.triggerEventHandler(value, String.valueOf(position));
                        if (selectedByCode) {
                            selectedByCode = false;
                        }else {
                            Fox.triggerFunction(parentController, value, String.valueOf(position), "", "");
                        }
                    }

                    @Override
                    public void onNothingSelected(AdapterView<?> parent) {

                    }
                });
                break;
            case "Enabled":
                v.setEnabled(value.equals("true"));
                break;
        }
        return "";
    }
}
