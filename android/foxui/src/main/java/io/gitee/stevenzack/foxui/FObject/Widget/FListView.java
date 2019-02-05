package io.gitee.stevenzack.foxui.FObject.Widget;

import android.support.v4.view.ViewCompat;
import android.support.v7.widget.RecyclerView;
import android.view.View;
import android.view.ViewGroup;

import org.json.JSONObject;
import org.json.JSONTokener;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

public class FListView extends FObject {
    public RecyclerView v;
    public MyAdapter adapter;
    public String fnOnGetView,fnOnBindData,fnOnGetCount;
    public String onItemClick,onItemLongClick;

    public FListView(FoxActivity foxActivity, RecyclerView.LayoutManager layoutManager) {
        parentController = foxActivity;
        v = new RecyclerView(parentController);
        v.setId(ViewCompat.generateViewId());
        view=v;
        v.setLayoutManager(layoutManager);
        adapter = new MyAdapter(this);
        v.setAdapter(adapter);
    }

    @Override
    public String getAttr(String attr) {
        String str = super.getAttr(attr);
        if (str!= null) {
            return str;
        }
        switch (attr) {

        }
        return "";
    }

    @Override
    public String setAttr(String attr, String value, String value2) {
        String str = super.setAttr(attr, value, value2);
        if (str != null) {
            return str;
        }
        switch (attr) {
            case "OnGetView":
                fnOnGetView=value;
                break;
            case "OnBindData":
                fnOnBindData=value;
                break;
            case "OnGetCount":
                fnOnGetCount = value;
                break;
            case "NotifyDataSetChanged":
                v.getAdapter().notifyDataSetChanged();
                break;
            case "OnItemClick":
                onItemClick=value;
                break;
            case "OnItemLongClick":
                onItemLongClick=value;
                break;
        }
        return "";
    }

    class MyAdapter extends RecyclerView.Adapter<MyAdapter.ViewHolder> {
        private FListView parentListView;
        private OnItemClickListener mItemClickListener=new OnItemClickListener() {
            @Override
            public void onItemClick(View view) {
                if (parentListView.onItemClick==null||parentListView.onItemClick.equals(""))
                    return;
                int pos=parentListView.v.getChildAdapterPosition(view);
//                Fox.triggerEventHandler(parentListView.onItemClick,String.valueOf(pos));
                Fox.triggerFunction(parentController,onItemClick,String.valueOf(pos),"","");
            }

            @Override
            public void onItemLongClick(View view) {
                if (parentListView.onItemLongClick==null||parentListView.onItemLongClick.equals(""))
                    return;
                int pos=parentListView.v.getChildAdapterPosition(view);
//                Faithdroid.triggerEventHandler(parentListView.onItemLongClick,String.valueOf(pos));
                Fox.triggerFunction(parentController, onItemLongClick, String.valueOf(pos), "", "");
            }
        };

        public class ViewHolder extends RecyclerView.ViewHolder {
            public FObject fView;
            public String str;
            public ViewHolder(FObject v,String str) {
                super(v.view);
                fView = v;
                this.str =str;
            }
        }
        public MyAdapter(FListView fListView) {
            parentListView =fListView;
        }
        @Override
        public MyAdapter.ViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
            if (parentListView.fnOnGetView == null) {
                return new ViewHolder(null,"");
            }
            String viewGot= Fox.triggerFunction(parentController,parentListView.fnOnGetView,"","","");

            try {
                JSONObject object = (JSONObject) (new JSONTokener(viewGot).nextValue());
                FObject v = parentListView.parentController.viewmap.get(object.getString("VID"));
                v.view.setLayoutParams(FFrameLayout.parseLP(v));
                MyAdapter.ViewHolder vh = new ViewHolder(v, viewGot);
                vh.itemView.setOnClickListener(new View.OnClickListener() {
                    @Override
                    public void onClick(View v) {
                        mItemClickListener.onItemClick(v);
                    }
                });
                vh.itemView.setOnLongClickListener(new View.OnLongClickListener() {
                    @Override
                    public boolean onLongClick(View v) {
                        mItemClickListener.onItemLongClick(v);
                        return true;
                    }
                });
                return vh;
            } catch (Exception e) {
                e.printStackTrace();
                return null;
            }
        }
        @Override
        public void onBindViewHolder(ViewHolder holder, int position) {
            if (parentListView.fnOnBindData == null) {
                return;
            }
            try {
                JSONObject jo = new JSONObject();
                jo.put("Str", holder.str);
                jo.put("Position", position);
                Fox.triggerFunction(parentController,parentListView.fnOnBindData, jo.toString(),"","");
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
        @Override
        public int getItemCount() {
            if (parentListView.fnOnGetCount == null) {
                return 0;
            }
            try {
                return Integer.parseInt(Fox.triggerFunction(parentController,parentListView.fnOnGetCount,"","",""));
            } catch (Exception e) {
                e.printStackTrace();
            }
            return 0;
        }
    }
    public static interface OnItemClickListener {
        void onItemClick(View view);
        void onItemLongClick(View view);
    }
}
