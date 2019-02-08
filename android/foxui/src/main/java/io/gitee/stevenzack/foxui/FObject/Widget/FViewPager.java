package io.gitee.stevenzack.foxui.FObject.Widget;

import android.os.Bundle;
import android.support.v4.app.Fragment;
import android.support.v4.app.FragmentManager;
import android.support.v4.app.FragmentStatePagerAdapter;
import android.support.v4.view.ViewCompat;
import android.support.v4.view.ViewPager;

import org.json.JSONArray;
import org.json.JSONObject;
import org.json.JSONTokener;

import java.util.ArrayList;
import java.util.List;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

public class FViewPager extends FObject  {
    private FaithCollectionPagerAdapter adapter;
    public ViewPager v;
    public List<FPage> pages = new ArrayList<>();
    public String onCreateView,onGetCount,onPageSelected;
    public FTabLayout bindFTabLayout;

    public FViewPager(FoxActivity activity) {
        parentController = activity;
        v = new ViewPager(parentController);
        v.setId(ViewCompat.generateViewId());
        v.addOnPageChangeListener(new ViewPager.OnPageChangeListener() {
            @Override
            public void onPageScrolled(int i, float v, int i1) {

            }

            @Override
            public void onPageSelected(int i) {
                if (onPageSelected != null) {
                    Fox.triggerFunction(parentController,onPageSelected, String.valueOf(i),"","");
                }
            }

            @Override
            public void onPageScrollStateChanged(int i) {

            }
        });
        view=v;
        parseSize("-1","-1");
    }

    @Override
    public String getAttr(String attr) {
        String str = super.getAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr){
            // --------------------------------------------------
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
            case "Pages":
                try {
                    JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    for(int i=0;i<array.length();i++){
                        JSONObject object = array.getJSONObject(i);
                        FPage fPage=new FPage();
                        fPage.vid = object.getString("VID");
                        pages.add(fPage);
                    }
                    if(adapter==null) {
                        adapter = new FaithCollectionPagerAdapter(parentController.getSupportFragmentManager(), this);
                        v.setAdapter(adapter);
                    }
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "TabLayout":
                FObject o = parentController.viewmap.get(value);
                if (o==null)break;
                FTabLayout fTabLayout=(FTabLayout)o;
                bindFTabLayout=fTabLayout;
                fTabLayout.v.setupWithViewPager(v);
                break;
            case "OnCreateView":
                onCreateView = value;
                if (adapter==null&&onGetCount!=null) {
                    adapter = new FaithCollectionPagerAdapter(parentController.getSupportFragmentManager(), this);
                    v.setAdapter(adapter);
                }
                break;
            case "OnGetCount":
                onGetCount=value;
                if (adapter==null&&onCreateView!=null) {
                    adapter = new FaithCollectionPagerAdapter(parentController.getSupportFragmentManager(), this);
                    v.setAdapter(adapter);
                }
                break;
            case "CurrentItem":
                int index = Integer.parseInt(value);
                boolean is = value2.equals("true");
                v.setCurrentItem(index,is);
                break;
            case "OnPageSelected":
                onPageSelected = value;
                break;
        }
        return "";
    }

    // ------------------
    public class FPage{
        public String vid;
    }
    class FaithCollectionPagerAdapter extends FragmentStatePagerAdapter {
        private final FViewPager fviewPager;

        public FaithCollectionPagerAdapter(FragmentManager fm, FViewPager fViewPager) {
            super(fm);
            this.fviewPager =fViewPager;
        }

        @Override
        public Fragment getItem(int i) {
            Fragment fragment = new DemoObjectFragment().setFviewPager(fviewPager);
            Bundle args = new Bundle();
            // Our object is just an integer :-P
            args.putInt(DemoObjectFragment.ARG_OBJECT, i);
            fragment.setArguments(args);
            return fragment;
        }
        @Override
        public int getCount() {
            if (fviewPager.pages.size()==0&&onGetCount!=null&&onCreateView!=null) {
                try {
                    return Integer.parseInt(Fox.triggerFunction(parentController,onGetCount, "","",""));
                } catch (Exception e) {
                }
            }
            return fviewPager.pages.size();
        }
        @Override
        public CharSequence getPageTitle(int position) {
            if (bindFTabLayout==null||bindFTabLayout.tabsList.size()<=position){
                return "Page "+position;
            }
            return bindFTabLayout.tabsList.get(position);
        }
    }
}
